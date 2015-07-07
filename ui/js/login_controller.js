(function(){
  angular.module('yggdrasil').controller('LoginController', LoginController);

  function LoginController($http, $cookies, $location, $mdToast) {
    var vm = this;

    function toast(msg) {
      $mdToast.show(
        $mdToast.simple()
          .content(msg)
          .position("top right")
          .hideDelay(3000)
      );
    }

    vm.login = function() {
      $http.post("/api/login", {username: vm.username, password: vm.password}).success(function(data) {
        $cookies.put("apiKey", data.apiKey);
        $location.path("/app");
      }).error(function(data, status) {
        if(status === 404) {
          toast("Username or password is incorrect");
        } else {
          toast("There was an error logging you in");
        }
      });
    };
  }
})();
