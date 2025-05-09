import javafx.collections.ObservableList;
import javafx.scene.shape.Polygon;

import java.io.Serializable;
import java.util.ArrayList;
import java.util.Arrays;

public class EPolygon extends Polygon implements Serializable {

    EPolygon(double[] array){
        super(array);
    }

void moveCenter(double newX, double newY) {
        ObservableList<Double> points = this.getPoints();
        double centerX;
        double centerY;
        centerX = points.stream().filter(aDouble -> points.indexOf(aDouble) % 2 == 0).mapToDouble(aDouble -> aDouble).average().orElse(0);
        centerY = points.stream().filter(aDouble -> points.indexOf(aDouble) % 2 == 1).mapToDouble(aDouble -> aDouble).average().orElse(0);
        double deltaX = newX - centerX;
        double deltaY = newY - centerY;
        ArrayList<Double> newPoints = new ArrayList<>();

        for (int i = 0; i < points.size(); i++) {
            if (i % 2 == 0) {
                newPoints.add(points.get(i) + deltaX);
            } else {
                newPoints.add(points.get(i) + deltaY);
            }
        }

        this.getPoints().clear();
        this.getPoints().addAll(newPoints);

}
}
