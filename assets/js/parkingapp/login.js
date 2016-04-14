$(document).keypress(function(e) {
    if (e.which == 13) {
        doLogin();
    }
});

function doLogin() {
  user_name = $("#username").val();
  pass = $("#password").val();

  var form_data = {
      username: user_name,
      password: pass
  };

  if (user_name ==='' ||pass === '')
  {
      alert ("Username or Password Can't be empty");
      exit();
  }

  else{
      //alert("Masuk sini");
      $.post("api/user/login",
      {
          username:user_name,
          password:pass,
      },
  function(data,status){
          //alert("Data: " + data + "\nStatus: " + status);
          if (data.Status) {
              window.location.replace("/user/user_auth");
          } else {
              alert(data.Message);
          }

      });
    }
}
