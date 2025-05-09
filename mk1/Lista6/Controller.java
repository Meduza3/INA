import javafx.fxml.FXML;
import javafx.fxml.Initializable;
import javafx.scene.Node;
import javafx.scene.layout.Pane;
import javafx.scene.paint.Color;

import java.net.URL;
import java.util.ResourceBundle;

public class Controller implements Initializable {
    //PARAMETERS
    private double cellwidth;
    private double cellHeight;
    private static int speed = 1000;
    private static double probability = 0.2;
    static int n;
    static int m;
    @FXML
    public Pane pane;

    public RandomGenerator rng = new RandomGenerator();

    //METHODS
    @Override
    public void initialize(URL location, ResourceBundle resources) {
        System.out.println("Initializing");
        n = 8; //Number of cells in a row
        m = 8; //Number of cells in a column
        cellwidth = (double) 1000 / n;
        cellHeight = (double) 500 / m;
        for (int i = 0; i < n; i++) {
            for (int j = 0; j < m; j++) {
                Cell cell = new Cell(i * cellwidth, j * cellHeight, cellwidth, cellHeight);
                cell.setId(Integer.toString(i) + "," + Integer.toString(j));
                cell.setOnMouseClicked(event -> {
                    System.out.println("Clicked on " + cell.getId());
                    if(cell.isPaused){
                        cell.isPaused = false;
                        System.out.println("Unpausing");
                    } else {
                        cell.isPaused = true;
                        System.out.println("Pausing");
                    }
                });
                Color col = new Color(rng.nextDouble(), rng.nextDouble(), rng.nextDouble(), 1);
                cell.setFill(col);
                //
                cell.setStroke(col);
                pane.getChildren().add(cell);
                System.out.println("Added a cell at " + i * cellwidth + ", " + j * cellHeight);
            }
        }
        System.out.println("Initialized");
        startThreads();
    }

    private void startThreads(){
        for (Node node : pane.getChildren()) {
            Thread thread = new Thread((Cell) node, node.getId());
            thread.start();
        }
    }
    public static double getProbability() {
        return probability;
    }
    static public int getSpeed(){
        return speed;
    }

    public static int getn(){
        return n;
    }

    public static int getm(){
        return m;
    }
}
