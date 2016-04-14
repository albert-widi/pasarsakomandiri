$("#check_ticket").on('click', function() {
  var ticketNumber = document.getElementById("ticket_number").value;

  $.get("/api/parking/getTicketInfo", {
        ticket_number:ticketNumber
    },function(data, status){
      if (data.Response.Status == "Failed") {
        alert(data.Response.Message);
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


        $.get("/api/ipcamera/getPictureFromDevice", {
            ticket_number:ticketNumber
        },function(data, status){
            //alert(data.Data.Id);
            document.getElementById("picture_out_id").value = data.Data.Id;
            $("#picture_out").attr("src",data.Data.PictureFullPath);
        });
    });
});
