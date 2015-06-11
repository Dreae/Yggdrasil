package daemon

import (
  "os"
  "io"
  "bytes"
  "strconv"
  "net/http"
  "io/ioutil"
  "archive/zip"
  "path/filepath"
)

func Init_SteamCmd() error {
  _, err := os.Stat(".yggdrasil/steamcmd.exe")

  if err != nil && os.IsNotExist(err) {
    resp, err := http.Get("http://media.steampowered.com/installer/steamcmd.zip")
    if err != nil {
      return err
    }
    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
      return err
    }

    size, _ := strconv.ParseInt(resp.Header.Get("Content-Length"), 10, 32)
    reader, err := zip.NewReader(bytes.NewReader(body), size)

    for _, f := range reader.File {
      err := extractAndWrite(f)
      if err != nil {
        return err
      }
    }
    if err != nil {
      return err
    }

  } else {
    return err
  }
  return nil
}

func extractAndWrite(f *zip.File) error {
  rc, err := f.Open()
  if err != nil {
    return err
  }
  defer rc.Close()

  path := filepath.Join(".yggdrasil", f.Name)
  if f.FileInfo().IsDir() {
    os.MkdirAll(path, f.Mode())
  } else {
    f, err := os.OpenFile(path, os.O_WRONLY | os.O_CREATE | os.O_TRUNC, f.Mode())
    if err != nil {
      return err
    }
    defer f.Close()

    _, err = io.Copy(f, rc)
    if err != nil {
      return err
    }
  }
  return nil
}
