class Transmitter {

  String name;
  int position;
  int attemptCounter = 1;
  int signalPositionL;
  int signalPositionR;
  boolean transmitting;
  int delay;
  boolean msgSent;
  String color;
  int waitTime = 0;
  int transmissionFailCounter;


  Transmitter(String name, int position, String color){
    this.name = name;
    this.position = position;
    signalPositionR = position + 1;
    signalPositionL = position - 1;
    transmitting = false;
    this.color = color;
  }
}