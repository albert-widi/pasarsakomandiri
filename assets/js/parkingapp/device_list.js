$(document).ready(function(){
    dType();
    dName();
    buttonFunc();
});

function dType(){
    $.get("/formval",function(data, status){

        //alert("Data: " + data + "\nStatus: " + status);
        var kluar = document.getElementById("type");
        var ica = "";
             for (var i=0; i<data.length; i++) {
                //alert("oke gan");
                //alert(data[i].Device_type;
                ica += "<option value="+data[i].Device_type+">" + data[i].Device_type + "</option>";
               }
        kluar.innerHTML = ica;

    });
}

function dName(){
    $.get("/api/device/device_list", function(data, status){

          var div = document.getElementById("out");
          var odong = "";
          //alert("Data length: " + data.length)
            for (var i=0; i<data.length; i++) {

                id = data[i].Id;
                odong += "<tr><td>" + data[i].Device_type + "</td><td>" + data[i].Device_name + "</td><td>" + data[i].Host + "</td><td>" + data[i].Token + "</td><td>" + data[i].Description + "</td><td><button class='waves-effect waves-light btn' onclick='editFunc("+id+")'>Edit</button></td><td><button class='waves-effect waves-light btn' onclick='passFunc("+id+")'>Delete</button></td></tr>";
            }
            odong += ""
            div.innerHTML = odong;

    });
}

function buttonFunc(){
    $("button").click(function(){

       if (confirm("Anda Yakin???") == true) {
              Utype = document.getElementById("type").value;//$("type").val();
              Uname = $("#name").val();
              Uhost = $("#host").val();
              Utoken = $("#token").val();
              Udescription = $("#description").val();

                   $.post("/api/device/create_device",
                        {
                           type: Utype,
                           name: Uname,
                           host: Uhost,
                           token: Utoken,
                           description: Udescription
                        },
                       function(data,status){
                        //alert("Data: " + data + "\nStatus: " + status);
                        alert("Status" + data.Status + ", Message: " + data.Message);

                            if(data.Status === 'Success') {
                                $.get("/api/device/device_list", function(data, status) {

                                     var div = document.getElementById("out");
                                     var odong = "<table>";
                                     //alert("Data length: " + data.length)
                                       for (var i=0; i<data.length; i++) {

                                           id = data[i].Id;
                                           odong += "<tr><td>" + data[i].Device_type + "</td><td>" + data[i].Device_name + "</td><td>" + data[i].Host + "</td><td>" + data[i].Token + "</td><td>" + data[i].Description + "</td><td><button class='waves-effect waves-light btn' onclick='editFunc("+id+")'>Edit</button></td><td><button class='waves-effect waves-light btn' onclick='passFunc("+id+")'>Delete</button></td></tr>";
                                       }
                                       odong += "</table>"
                                       div.innerHTML = odong;

                                      });

                             }
                        });
           } else {

           }
    });
}