angular.module('yggdrasil', ['ngMaterial', 'ui.router', 'ngCookies']);

angular.module('yggdrasil').config(function($stateProvider, $urlRouterProvider){
  $urlRouterProvider.otherwise("/app");

  $stateProvider.state("login", {
    url: "/login",
    templateUrl: "partials/login.html",
    controller: "LoginController as loginCtrl"
  }).state("app", {
    url: "/app",
    templateUrl: "partials/app.html",
    controller: "AppController as app"
  });
});
