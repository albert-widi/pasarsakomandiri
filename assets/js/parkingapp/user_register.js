$(document).ready(function(){
    getTable();
    getRoleList();
    buttonCreate();
    sideNav();
});

function getTable(){
    $.get("/api/user/all_user", function(data,status){

        var div = document.getElementById ("display");
        var tables = "";
        for (var i=0; i<data.length; i++){
            var id = data[i].Id;
            tables += "<tr><td>" + data[i].Username + "</td> <td>" + data[i].Role + "</td> <td>" + data[i].Description +  "</td><td><button class='waves-effect waves-light btn' onclick='passFunc("+id+")'>Edit</button></td></tr>";
        }
        tables += ""
        div.innerHTML = tables;

    });
}

function getRoleList(){
    $.get ("/api/user/role_list", function (data, status){
        var div = document.getElementById("slctRole");
        var options = "";
        for (i=0; i<data.length; i++){
            options += "<option value= "+data[i].Role_name+">" +data[i].Role_name+ "</option>";
        }
        div.innerHTML = options;
    });
}

function buttonCreate(){
    $("#create").click(function(){
        username = $("#username").val();
        password = $("#password").val();
        role = document.getElementById("slctRole").value;
        desc = $("#description").val();

            $.post("/api/user/register",
                {
                    username: username,
                    password: password,
                    role: role,
                    description: desc
                },

                function(data, status) {
                    alert(data.Message);

                    if(data.Status == "Success") {
                        $.get("/api/user/all_user", function(data,status){
                            var div = document.getElementById ("display");
                            var tables = "";
                            for (var i=0; i<data.length; i++){
                                var id = data[i].Id;
                                tables += "<tr><td>" + data[i].Username + "</td> <td>" + data[i].Role + "</td> <td>" + data[i].Description +  "</td><td><button class='waves-effect waves-light btn' onclick='passFunc("+id+")'>Edit</button></td></tr>";
                            }
                            tables += ""
                            div.innerHTML = tables;
                        });
                    }
        });
    });
}

function sideNav(){
    $(function(){
         $('.button-collapse').sideNav({
               menuWidth: 240, // Default is 240
               edge: 'left', // Choose the horizontal origin
               closeOnClick: false // Closes side-nav on <a> clicks, useful for Angular/Meteor
             }
           );
    });
}