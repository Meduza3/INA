import javafx.event.EventHandler;
import javafx.scene.input.MouseEvent;
import javafx.scene.paint.Color;
import javafx.scene.shape.Circle;

import java.io.Serializable;

public class ECircle extends Circle implements Serializable {
    private double x;
    private double y;
    private double r;
    private double red;
    private double blue;
    private double green;
    private double scaleX, scaleY, translateX, translateY;

    ECircle(double x, double y, double r){
        this.x = x;
        this.y = y;
        this.r = r;
    }
    void moveCenter(double newX, double newY){
        this.setCenterX(newX);
        this.setCenterY(newY);
    }

    void resize(double newR){
        this.setRadius(newR);
    }


    public void saveProperties() {
        x = this.getCenterX();
        y = this.getCenterY();
        r = this.getRadius();
        red = ((Color)this.getFill()).getRed();
        green = ((Color)this.getFill()).getGreen();
        blue = ((Color)this.getFill()).getBlue();
        scaleX = this.getScaleX();
        scaleY = this.getScaleY();
        translateX = this.getTranslateX();
        translateY = this.getTranslateY();
    }

    public void setProperties() {
        this.setCenterX(x);
        this.setCenterY(y);
        this.setRadius(r);
        this.setScaleX(scaleX);
        this.setScaleY(scaleY);
        this.setFill(Color.color(red, green, blue));
        this.setTranslateX(translateX);
        this.setTranslateY(translateY);
    }
}
