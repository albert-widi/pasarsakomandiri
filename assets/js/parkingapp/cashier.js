


//focus ticket number when 1st load
$(document).ready(function(){
    $("#ticket_number").focus();
    //window.flagenter = 0;
});

var flagenter = 0;



//trigger enter for check and input
$(document).keypress(function(e){

    if (e.which == 13){

       flagenter += 1;

        if (flagenter % 2 === 0){
            checkOut();
            flagenter = 0;
            $("#ticket_number").focus();
            exit();
        }

       var check = checkTicket();

    }
});


//click function
$(document).ready(function(){
    $("#check_ticket").click(function(){
        checkTicket();
    });

    $("#check_out").click(function(){
            checkOut();
        });
});



//function
function checkTicket(){
  var ticketNumber = document.getElementById("ticket_number").value;
  $.get("/api/parking/getTicketInfo", {
        ticket_number:ticketNumber
    },function(data, status){
      if (data.Response.Status == "Failed") {
        alert(data.Response.Message);
        flagenter = 0;
        $("#ticket_number").val('');
        return;
      }

       var labelTicketNumber = document.getElementById("label_ticket_number");
       document.getElementById("ticketnum").value = data.Data.Ticket_number;
       //labelTicketNumber.innerHTML = data.Data.Ticket_number;
       var labelDateIn = document.getElementById("label_date_in");
       document.getElementById("datein").value = data.Data.Created_date;
       //labelDateIn.innerHTML = data.Data.Created_date;
       var labelDateOut = document.getElementById("label_date_out");
       document.getElementById("dateout").value = data.Data.Out_date;
       //labelDateOut.innerHTML = data.Data.Out_date;
       var labelDuration = document.getElementById("label_duration");
       document.getElementById("duration").value = data.DeltaTime;
       //labelDuration.innerHTML = data.DeltaTime;
       var labelParkingCost = document.getElementById("label_parking_cost");
       document.getElementById("cost").value = data.Data.Parking_cost;
       //labelParkingCost.innerHTML = data.Data.Parking_cost;
       document.getElementById("ticket_id").value = data.Data.Id;

       if(data.Data.Ticket_number !== ''){
           $('#check_out').removeClass("disabled");
       }
        $("#picture_in").attr("src",data.Picture_path_in);
        $("#check_out").removeClass("disabled");
        $("#no_kendaraan").focus();

        $.get("/api/ipcamera/getPictureFromDevice", {
            ticket_number:ticketNumber
        },function(data, status){
            document.getElementById("picture_out_id").value = data.Data.Id;
            $("#picture_out").attr("src",data.Data.PictureFullPath);
        });

    });
}


function checkOut(){
    var ticketId = document.getElementById("ticket_id").value;
                    var ticketNumber = document.getElementById("ticket_number").value;
                    var vehicleNo = document.getElementById("no_kendaraan").value;
                    var dateOut = document.getElementById("dateout").value;
                    var cost = document.getElementById("cost").value;
                    var pictureOutId = document.getElementById("picture_out_id").value;
    		        var duration = document.getElementById("duration").value;

                    $.post("/api/parking/checkOut", {
                                    ticket_number:ticketNumber,
                                    ticket_id:ticketId,
                                    vehicle_number:vehicleNo,
                                    ticket_date_out:dateOut,
                                    parking_cost:cost,
                                    picture_out_id:pictureOutId,
    				parking_duration:duration
                                },function(data, status){

                                    $("#ticket_number").val('');
                                    $("#no_kendaraan").val('');
                                    $("#datein").val('');
                                    $("#dateout").val('');
                                    $("#duration").val('');
                                    $("#cost").val('');
                                    $("#ticketnum").val('');
                                    $("#picture_in").attr("src","nothing");
                                    $("#picture_out").attr("src","nothing");
                                    $("#check_out").addClass("disabled");


                                  if (data.Response.Status == "Failed") {
                                    alert(data.Response.Message);
                                    return;
                                  }

                                });
}
