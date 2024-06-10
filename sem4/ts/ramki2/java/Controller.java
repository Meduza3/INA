

import java.util.ArrayList;
import java.util.Random;


class Controller {


  private static final String ANSI_RED = "\u001B[31m";
  private static final String ANSI_RESET = "\u001B[0m";

  private String[] network;
  private ArrayList<Transmitter> transmitterList;
  private boolean arrayNoised = false;
  private Random r = new Random();
  private int delayCounter;
  private int printSpeed;
  private int delayRange;
  private boolean changeWaitTime = false;

  Controller(String[] network, ArrayList<Transmitter> transmitterList, int printSpeed, int delayRange) {
    this.network = network;
    this.transmitterList = transmitterList;
    this.printSpeed = printSpeed;
    this.delayRange = delayRange;
    clearNetwork();
  }

  private boolean checkIfpossibleTranssmition(Transmitter transmitter) {
    if (network[transmitter.position].equals("0")) {
      return true;
    }
    return false;
  }


  void run() throws InterruptedException {


    delayCounter = 1;
    randomizeDelay(transmitterList);

    while (true) {

      for (Transmitter transmitter : transmitterList) {

        boolean failTransmission = false;


        if (transmitter.delay <= delayCounter) {
          if (arrayNoised) {
            startClearing(transmitter);
          } else if (transmitter.msgSent) {
            transmitter.attemptCounter = 1;
            startClearingMSG(transmitter);
          } else if (checkIfArrayNoised()) {
            arrayNoised = true;
            changeWaitTime = true;
          } else if (transmitter.transmitting) {
            if (transmitter.waitTime < 2) {
              continueTransmission(transmitter);
            } else {
              transmitter.waitTime--;
              transmitter.transmitting = false;
            }
          } else if (checkIfpossibleTranssmition(transmitter)) {
            transmitter.transmitting = true;
            transmitter.waitTime--;
          } else if (isOtherTransmitting(transmitter)) {
            if(transmitter.waitTime < 2) {
              failTransmission = true;
            } else {
              transmitter.waitTime--;
            }
          }
        }


        if (failTransmission) {
          waitTimeAdd(transmitter);
        } else if (changeWaitTime) {
          changeWaitTime = false;
          for (Transmitter trans : transmitterList) {
            if(trans.waitTime<2)
            waitTimeAdd(trans);
          }
        }
      }


      printNetwork();
      Thread.sleep(printSpeed);
      delayCounter++;
    }
  }

  private boolean isOtherTransmitting(Transmitter transmitter) {
    if (network[transmitter.position].equals("=") || network[transmitter.position].equals("0")) {
      return false;
    }
    return true;
  }

  private void waitTimeAdd(Transmitter transmitter) {
    transmitter.attemptCounter = transmitter.attemptCounter * 2;
    if (transmitter.attemptCounter > 1024) {
      if (transmitter.transmissionFailCounter == 6) {
        System.out.println("TRANSMISSON " + transmitter.name + " FAILED");
      }
      transmitter.transmissionFailCounter++;
    }
    transmitter.waitTime = r.nextInt(transmitter.attemptCounter) + 1;
    transmitter.waitTime = transmitter.waitTime * network.length;
  }

  private void randomizeDelay(ArrayList<Transmitter> transmitters) {

    for (Transmitter trans : transmitters) {
      trans.delay = r.nextInt(delayRange);
    }
  }

  private boolean checkIfArrayNoised() {

    for (String aNetwork : network) {
      if (!aNetwork.equals("=")) {
        return false;
      }
    }
    return true;
  }


  private void startClearingMSG(Transmitter transmitter) {

    for(Transmitter trans : transmitterList){
      trans.transmitting = false;
    }

    boolean if1 = false;
    boolean if2 = false;

    if (transmitter.position > network.length / 2) {
      if1 = true;
    } else {
      if2 = true;
    }

    if (transmitter.signalPositionR < network.length) {
      if (if2) {
        network[transmitter.position] = "0";
      }
      network[transmitter.signalPositionR] = "0";
      transmitter.signalPositionR++;
    }

    if (transmitter.signalPositionL >= 0) {
      if (if1) {
        network[transmitter.position] = "0";
      }
      network[transmitter.signalPositionL] = "0";
      transmitter.signalPositionL--;
    }


    if (transmitter.signalPositionL == -1 && transmitter.signalPositionR == network.length) {
      transmitter.msgSent = false;
      transmitter.signalPositionR = transmitter.position + 1;
      transmitter.signalPositionL = transmitter.position - 1;
      transmitter.delay = r.nextInt(delayRange) + delayCounter;
    }

    if (arrayClear()) {
      transmitter.msgSent = false;
    }
  }

  private void startClearing(Transmitter transmitter) {

    boolean if1 = false;
    boolean if2 = false;

    if (transmitter.position > network.length / 2) {
      if1 = true;
    } else {
      if2 = true;
    }


    for(Transmitter trans : transmitterList){
      trans.transmitting = false;
    }


    if (transmitter.signalPositionR < network.length) {
      if (if2) {
        network[transmitter.position] = "0";
      }
      network[transmitter.signalPositionR] = "0";
      transmitter.signalPositionR++;
    }

    if (transmitter.signalPositionL >= 0) {
      if (if1) {
        network[transmitter.position] = "0";
      }
      network[transmitter.signalPositionL] = "0";
      transmitter.signalPositionL--;
    }

    if (transmitter.signalPositionL == -1 && transmitter.signalPositionR == network.length) {
      transmitter.msgSent = false;
    }

    if (arrayClear()) {

      arrayNoised = false;
      if (!transmitter.msgSent) {
        transmitter.delay = r.nextInt(delayRange) + delayCounter;
      }
      for (Transmitter trans : transmitterList) {
        trans.signalPositionR = trans.position + 1;
        trans.signalPositionL = trans.position - 1;
        transmitter.msgSent = false;
      }
    }
  }

  private boolean arrayClear() {
    for (String str : network) {
      if (!str.equals("0")) {
        return false;
      }
    }
    return true;
  }

  private void continueTransmission(Transmitter transmitter) {
    if (transmitter.transmitting && network[transmitter.position].equals("0")) {
      network[transmitter.position] = transmitter.name;
    }

    if (transmitter.signalPositionR < network.length) {
      sendMsgRight(transmitter);
    }
    if (transmitter.signalPositionL >= 0) {
      sendMsgLeft(transmitter);
    }


    if (transmitter.signalPositionL == -1 && transmitter.signalPositionR == network.length) {
      transmitter.transmitting = false;
      transmitter.signalPositionL = transmitter.position - 1;
      transmitter.signalPositionR = transmitter.position + 1;
    }
    transmitter.msgSent = checkIfMsgSent(transmitter);
  }

  private void sendMsgRight(Transmitter transmitter) {


    if (!network[transmitter.signalPositionR].equals("0") && !network[transmitter.signalPositionR].equals(transmitter.name)) {
      network[transmitter.signalPositionR] = "=";
    } else {
      network[transmitter.signalPositionR] = transmitter.name;
    }
    transmitter.signalPositionR++;
  }

  private void sendMsgLeft(Transmitter transmitter) {


    if (!network[transmitter.signalPositionL].equals("0") && !network[transmitter.signalPositionL].equals(transmitter.name)) {
      network[transmitter.signalPositionL] = "=";
    } else {
      network[transmitter.signalPositionL] = transmitter.name;
    }
    transmitter.signalPositionL--;
  }

  private void printNetwork() {


    for (String aNetwork : network) {
      for (Transmitter transmitter : transmitterList) {
        if (aNetwork.equals(transmitter.name)) {
          System.out.print(transmitter.color + aNetwork + ANSI_RESET + " ");
        } else if (aNetwork.equals("0")) {
          System.out.print(aNetwork + " ");
          break;
        } else if (aNetwork.equals("=")) {
          System.out.print(ANSI_RED + aNetwork + ANSI_RESET + " ");
          break;
        }
      }
    }
    System.out.println();
  }

  private void clearNetwork() {
    for (int i = 0; i < network.length; i++) {
      network[i] = "0";
    }
  }

  private boolean checkIfMsgSent(Transmitter transmitter) {
    for (String str : network) {
      if (!str.equals(transmitter.name)) {
        return false;
      }
    }
    return true;
  }
}