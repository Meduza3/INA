import javafx.scene.paint.Color;
import javafx.scene.shape.*;

import java.io.Serializable;

public class ERectangle extends Rectangle implements Serializable {

    public double x;
    public double y;
    public double width;
    public double height;
    public Color fillColor;
    public Color strokeColor;
    public int rotationNumber;
    double red, blue, green;

    public ERectangle(double x, double y, double width, double height) {
        super(x, y, width, height);
    }

    public ERectangle(double x, double y, double width, double height, Color fillColor, Color strokeColor, int rotationNumber) {
        this.x = x;
        this.y = y;
        this.width = width;
        this.height = height;
        this.fillColor = fillColor;
        this.strokeColor = strokeColor;
        this.rotationNumber = rotationNumber;
    }

    void moveCenter(double newX, double newY){
        setX(newX - getWidth()/2);
        setY(newY - getHeight()/2);
    }

    private void resizeRect(double newW, double newH){
        setWidth(newW);
        setHeight(newH);
    }

    @Override
    public String toString() {
        return "Rectangle [" + x + "," + y + "," + width + "," + height + "," + fillColor + "," + strokeColor + "]";
    }
    public static ERectangle fromString(String s){
        String[] parts = s.substring("Rectangle [".length(), s.length() - 1).split(",");
        int x = Integer.parseInt(parts[0]);
        int y = Integer.parseInt(parts[1]);
        int width = Integer.parseInt(parts[2]);
        int height = Integer.parseInt(parts[3]);
        //Color stroke = Integer.parseInt(parts[4]);
        //Color fill = Integer.parseInt(parts[5]);
        //return new ERectangle(x, y, width, height, stroke, fill);
        return new ERectangle(x, y, width, height);
    }

    public void saveProperties(){

    }
}
