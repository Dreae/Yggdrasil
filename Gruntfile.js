module.exports = function(grunt) {

  // Project configuration.
  grunt.initConfig({
    pkg: grunt.file.readJSON('package.json'),
    wiredep: {
      main: {
        src: [
          "ui/index.html"
        ],
        options: {
          ignorePath: "../",
          fileTypes: {
            html: {
              replace: {
                js: '<script src="/lib/{{filePath}}"></script>',
                css: '<link rel="stylesheet" href="/lib/{{filePath}}" />'
              }
            }
          }
        }
      }
    },
    copy: {
        bower: {
            files: [
                {expand: true, src: ['bower_components/**/angular.js'], dest: 'ui/lib'},
                {expand: true, src: ['bower_components/angular-animate/angular-animate.js'], dest: 'ui/lib'},
                {expand: true, src: ['bower_components/angular-aria/angular-aria.js'], dest: 'ui/lib'},
                {expand: true, src: ['bower_components/angular-material/angular-material.*'], dest: 'ui/lib'}
            ]
        }
    }
  });

  require('load-grunt-tasks')(grunt);
  grunt.registerTask('default', ['copy', 'wiredep']);
};
