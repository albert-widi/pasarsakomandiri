$(document).ready(function(){
    tableOutFirst();
    rasberryTipe();
    cameraType();
    vehicleType();
    clickDropdown();
    buttonCreateGroup();
    changeRegisterButton();
    cashierAll();
});


function tableOutFirst(){
    $.get("/api/device/device_group_list", function(data, status){
        var div = document.getElementById ("display");
        var tables = "";

        for (var n=0; n<data.length; n++){
            var id = data[n].Id;
            tables += "<tr><td>" +data[n].Raspberry_id+ "</td> <td>" +data[n].Raspberry_ip+ "</td> <td>" +data[n].Group_type+ "</td> <td>" +data[n].Camera_id+ "</td> <td>" +data[n].Camera_ip+ "</td> <td>" +data[n].Vehicle_id+ "</td> <td>" +data[n].Vehicle_type+ "</td> <td>" +data[n].Group_name+ "</td><td><button class='waves-effect waves-light btn' onclick= 'passFunc("+id+")'>Delete</button></td></tr>";
        }

        tables += "";
        div.innerHTML = tables;
    });
}

function rasberryTipe(){
    $.get("/api/device/devices_by_type", {
        device_type:"Raspberry"
            },function(data, status){

            var raspberry = document.getElementById("device_raspberry");
           var content = "";
             for (var i=0; i<data.length; i++) {
                content += "<option value="+data[i].Id+">" + data[i].Device_name + "</option>";
                }
            raspberry.innerHTML = content;
    });
}

function cameraType(){
    $.get("/api/device/devices_by_type", {
        device_type:"Camera"
    },function(data, status){

            var raspberry = document.getElementById("device_camera");
           var content = "";
             for (var i=0; i<data.length; i++) {
                content += "<option value="+data[i].Id+">" + data[i].Device_name + "</option>";
                }
            raspberry.innerHTML = content;
    });
}


function vehicleType(){
    $.get("/api/parking/vehicleAll", {
        },function(data, status){
                //alert(data);
                var raspberry = document.getElementById("vehicle_type");
               var content = "";
                 for (var i=0; i<data.length; i++) {
                    content += "<option value="+data[i].Id+">" + data[i].Vehicle_type + "</option>";
                    }
                raspberry.innerHTML = content;
        });
}


function clickDropdown(){
    $("#device_raspberry").on('click', function() {
        var deviceId = document.getElementById("device_raspberry").value;
        $.get("/api/device/device_info", {
                        device_id:deviceId
                    },function(data, status){
                           var raspberry = document.getElementById("raspberry_info");
                           var content = "<tr><td><b>Host</b></td><td><b>Device Type</b></td><td><b>Description</b></td></tr><tr><td>" + data.Host + "</td><td>" + data.Device_type + "</td><td>" + data.Description + "</td></tr>";
                            raspberry.innerHTML = content;
                    });
    });

    $("#device_camera").on('click', function() {
        var deviceId = document.getElementById("device_camera").value;
        $.get("/api/device/device_info", {
                        device_id:deviceId
                    },function(data, status){
                           var raspberry = document.getElementById("camera_info");
                           var content = "<tr><td><b>Host</b></td><td><b>Device Type</b></td><td><b>Description</b></td></tr><tr><td>" + data.Host + "</td><td>" + data.Device_type + "</td><td>" + data.Description + "</td></tr>";
                            raspberry.innerHTML = content;
                    });
    });
}

function buttonCreateGroup(){
    $("#create_group").on('click', function() {
        var raspberryid = document.getElementById("device_raspberry").value;
        var cameraid = document.getElementById("device_camera").value;
        var vehicleid = document.getElementById("vehicle_type").value;
        var gatename = document.getElementById("gate_name").value;
        $.post("/api/device/create_device_group",
        {
            raspberry_id:raspberryid,
            camera_id:cameraid,
            vehicle_id:vehicleid,
            gate_name:gatename,
            group_type:"Gate"
        },function(data, status) {
            alert(data.Message);

            if(data.Status == 'Success'){
               $.get("/api/device/device_group_list", function(data, status){
                   var div = document.getElementById ("display");
                   var tables = "";

                   for (var n=0; n<data.length; n++){
                       var id = data[n].Id;
                       tables += "<tr><td>" +data[n].Raspberry_id+ "</td><td>" +data[n].Raspberry_ip+ "</td><td>" + "</td><td>" +data[n].Group_type+ "</td><td>" +data[n].Camera_id+ "</td> <td>" +data[n].Camera_ip+ "</td><td>" +data[n].Vehicle_id+ "</td> <td>" +data[n].Vehicle_type+ "</td><td>" +data[n].Group_name+ "</td><td><button class='waves-effect waves-light btn' onclick= 'passFunc("+id+")'>Delete</button></td></tr>";
                   }

                   tables += "";
                   div.innerHTML = tables;

               });
           }
        });
    });
}


function passFunc(id, buttonCreateGroup){
    alert (id);

    $.post("/api/device/delete_device_group",
    {
    id:id
    }, function (data){
        alert (data.Message);
        if(data.Status == 'Success'){
            $.get("/api/device/device_group_list", function(data, status){
               var div = document.getElementById ("display");
               var tables = "";

               for (var n=0; n<data.length; n++){
                   var id = data[n].Id;
                   tables += "<tr><td>" +data[n].Raspberry_id+ "</td><td>" +data[n].Raspberry_ip+ "</td> <td>" +data[n].Camera_id+ "</td> <td>" +data[n].Camera_ip+ "</td> <td>" +data[n].Vehicle_id+ "</td> <td>" +data[n].Vehicle_type+ "</td> <td>" +data[n].Gate_name+ "</td><td><button class='waves-effect waves-light btn' onclick= 'passFunc("+id+")'>Delete</button></td></tr>";
               }

               tables += "";
               div.innerHTML = tables;

            });
        }
    });

};


function changeRegisterButton(){
    $('#munc').on('click', function(){
                $('#ilang').removeClass('active');
                $('#munc').addClass('active');
                $('#devreg').css('display', 'none');
                $('#gatereg').css({'display': 'block', 'margin-top': '-5px'});
            });

            $('#ilang').on('click', function(){
                $('#munc').removeClass('active');
                $('#ilang').addClass('active');
                $('#gatereg').css('display', 'none');
                $('#devreg').css({'display': 'block', 'margin-top': '-5px'});
            });
}


function cashierAll(){
//get value cashier name
    $.get("/api/device/devices_by_type", {
        device_type:"Cashier"
            },function(data, status){

            var raspberry = document.getElementById("cashier_name");
           var content = "";
             for (var i=0; i<data.length; i++) {
                content += "<option value="+data[i].Id+">" + data[i].Device_name + "</option>";
                }
            raspberry.innerHTML = content;
    });

    //get value cashier camera
    $.get("/api/device/devices_by_type", {
        device_type:"Camera"
            },function(data, status){

            var raspberry = document.getElementById("cashier_camera");
           var content = "";
             for (var i=0; i<data.length; i++) {
                content += "<option value="+data[i].Id+">" + data[i].Device_name + "</option>";
                }
            raspberry.innerHTML = content;
    });

    //click func on cashier name
    $('#cashier_name').on('click', function(){

        var deviceId = document.getElementById("cashier_name").value;
        $.get("/api/device/device_info", {
                        device_id:deviceId
                    },function(data, status){
                           var raspberry = document.getElementById("cashier_info");
                           var content = "<tr><td><b>Host</b></td><td><b>Device Type</b></td><td><b>Description</b></td></tr><tr><td>" + data.Host + "</td><td>" + data.Device_type + "</td><td>" + data.Description + "</td></tr>";
                            raspberry.innerHTML = content;
         });

    });


    //click func on camera cashier name
    $('#cashier_camera').on('click', function(){

        var deviceId = document.getElementById("cashier_camera").value;
        $.get("/api/device/device_info", {
                        device_id:deviceId
                    },function(data, status){
                           var raspberry = document.getElementById("cashier_camera_info");
                           var content = "<tr><td><b>Host</b></td><td><b>Device Type</b></td><td><b>Description</b></td></tr><tr><td>" + data.Host + "</td><td>" + data.Device_type + "</td><td>" + data.Description + "</td></tr>";
                            raspberry.innerHTML = content;
         });

    });

    //create group button cashier
     $("#create_group_cashier").on('click', function() {
        var raspberryid = document.getElementById("cashier_name").value;
        var cameraid = document.getElementById("cashier_camera").value;
        var gatename = document.getElementById("gate_name_cashier").value;
        $.post("/api/device/create_device_group",
        {
            raspberry_id:raspberryid,
            camera_id:cameraid,
            gate_name:gatename,
            group_type:"Cashier"
        },function(data, status) {
            alert(data.Message);

            if(data.Status == 'Success'){
               $.get("/api/device/device_group_list", function(data, status){
                   var div = document.getElementById ("display");
                   var tables = "";

                   for (var n=0; n<data.length; n++){
                       var id = data[n].Id;
                       tables += "<tr><td>" +data[n].Raspberry_id+ "</td><td>" +data[n].Raspberry_ip+ "</td><td>" + "</td><td>" +data[n].Group_type+ "</td><td>" +data[n].Camera_id+ "</td> <td>" +data[n].Camera_ip+ "</td><td>" +data[n].Vehicle_id+ "</td> <td>" +data[n].Vehicle_type+ "</td><td>" +data[n].Group_name+ "</td><td><button class='waves-effect waves-light btn' onclick= 'passFunc("+id+")'>Delete</button></td></tr>";
                   }

                   tables += "";
                   div.innerHTML = tables;

               });
           }
        });
    });
}

//side-nav jquery
$(function(){
    $('.button-collapse').sideNav({
          menuWidth: 240, // Default is 240
          edge: 'left', // Choose the horizontal origin
          closeOnClick: true // Closes side-nav on <a> clicks, useful for Angular/Meteor
        });
});



