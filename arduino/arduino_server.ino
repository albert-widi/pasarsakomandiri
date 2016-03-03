/*
  This a simple example of the aREST Library for Arduino (Uno/Mega/Due/Teensy)
  using the Ethernet library (for example to be used with the Ethernet shield).
  See the README file for more details.

  Written in 2014 by Marco Schwartz under a GPL license.
*/

// Libraries
#include <SPI.h>
#include <Ethernet.h>
#include <aREST.h>
#include <avr/wdt.h>

// Enter a MAC address for your controller below.
byte mac[] = { 0x90, 0xA2, 0xDA, 0x0E, 0xFE, 0x40 };

// IP address in case DHCP fails
IPAddress ip(192,168,0,100);

// Ethernet server
EthernetServer server(80);
// Ethernet Client
EthernetClient newClient;

// Create aREST instance
aREST rest = aREST();

// Variables to be exposed to the API
int temperature;
int humidity;

//char testServer[] = "localhost";
IPAddress testServer(192,168,0,101);

void setup(void)
{
  // Start Serial
  Serial.begin(9660);

  // Init variables and expose them to REST API
  temperature = 24;
  humidity = 40;
  rest.variable("temperature",&temperature);
  rest.variable("humidity",&humidity);

  // Function to be exposed
  rest.function("led",ledControl);
  rest.function("test",testControl);
  rest.function("testpost", testPost);

  // Give name and ID to device
  rest.set_id("008");
  rest.set_name("dapper_drake");

  // Start the Ethernet connection and the server
  if (Ethernet.begin(mac) == 0) {
    Serial.println("Failed to configure Ethernet using DHCP");
    // no point in carrying on, so do nothing forevermore:
    // try to congifure using IP address instead of DHCP:
    Ethernet.begin(mac, ip);
  }
  server.begin();
  Serial.print("server is at ");
  Serial.println(Ethernet.localIP());

  // Start watchdog
  wdt_enable(WDTO_4S);
}

void loop() {

  // listen for incoming clients
  EthernetClient client = server.available();
  rest.handle(client);
  wdt_reset();

}

int testPost(String command) {
  Serial.println("Test Post");
  //delay(1000);
  if (newClient.connect(testServer, 8080)) {
    Serial.println("Connected");
    newClient.println("GET /opengate HTTP/1.1");
    newClient.println("Host: 192.168.0.101");
    newClient.println("Connection: close");
    newClient.println("");
    newClient.stop();
  } else {
    Serial.println("Not connected");
    newClient.stop();
  }
}

int testControl(String command) {
  Serial.println("Open The GATE!");
  return 1;
}

// Custom function accessible by the API
int ledControl(String command) {

  Serial.println(command);
  // Get state from command
  int state = command.toInt();

  digitalWrite(6,state);
  return 1;
}