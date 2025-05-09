import javafx.scene.layout.Pane;
import javafx.scene.paint.Color;
import javafx.scene.shape.Rectangle;
public class Cell extends Rectangle implements Runnable {
    private double x;
    private double y;
    private double width;
    private double height;
    private Color color;
    private String id;
    public RandomGenerator rng = new RandomGenerator();
    private boolean hasFinished = true;
    public boolean isPaused = false;
    public Cell(double x, double y, double width, double height) {
        super(x, y, width, height);

        setStroke(javafx.scene.paint.Color.BLACK);
        setFill(javafx.scene.paint.Color.WHITE);

        this.width = width;
        this.height = height;
        this.x = x;
        this.y = y;
    }


    @Override
    public void run() {
        try {
            Thread.sleep(10L * (rng.nextInt(Controller.getSpeed()) + Controller.getSpeed() / 2));
        } catch (InterruptedException e) {
            throw new RuntimeException(e);
        }
        while (true) {
            if (isPaused == false) {
                //System.out.println("Start: " + Thread.currentThread().getName());
                synchronized (this) { // synchronize on the current cell
                    while (!hasFinished) { // wait until the previous color computation has finished
                        try {
                            wait();
                        } catch (InterruptedException e) {
                            throw new RuntimeException(e);
                        }
                    }
                    // compute the new color
                    if (rng.nextDouble() < Controller.getProbability()) {
                        //System.out.println("Changing color");
                        Color col = getAverageColor();
                        setFill(col);
                        setStroke(col);
                    }
                    // set the flag to indicate that this color computation has finished
                    hasFinished = true;
                    notifyAll(); // notify other threads that this thread has finished
                }
                try {
                    Thread.sleep((long) (rng.nextInt(Controller.getSpeed()) + Controller.getSpeed() / 2));
                } catch (InterruptedException e) {
                    throw new RuntimeException(e);
                }
                //System.out.println("End: " + Thread.currentThread().getName());
            }
        }
    }
    private Cell getCell(int x, int y, Pane pane){
        return (Cell) pane.lookup("#"+Integer.toString(x) + "," + Integer.toString(y));
    }

    private Color getAverageColor(){
        Cell left = getCell((int) (x/width) - 1, (int) (y/height), (Pane) getParent());
        Cell right = getCell((int) (x/width) + 1, (int) (y/height), (Pane) getParent());
        Cell up = getCell((int) (x/width), (int) (y/height) - 1, (Pane) getParent());
        Cell down = getCell((int) (x/width), (int) (y/height) + 1, (Pane) getParent());

        if (left == null) {
            left = getCell(Controller.getn() - 1, (int) (y/height), (Pane) getParent());
        }
        if (up == null) {
            up = getCell((int) (x/width),Controller.getm() - 1, (Pane) getParent());
        }
        if (right == null) {
            right = getCell(0, (int) (y/height), (Pane) getParent());
        }
        if (down == null) {
            down = getCell((int) (x/width), 0, (Pane) getParent());
        }

        Color cl = (Color) left.getFill();
        Color cr = (Color) right.getFill();
        Color cu = (Color) up.getFill();
        Color cd = (Color) down.getFill();
        double red = (cl.getRed() + cr.getRed() + cu.getRed() + cd.getRed())/4;
        double green = (cl.getGreen() + cr.getGreen() + cu.getGreen() + cd.getGreen())/4;
        double blue = (cl.getBlue() + cr.getBlue() + cu.getBlue() + cd.getBlue())/4;
        return new Color(red, green, blue, 1);
    }
}
