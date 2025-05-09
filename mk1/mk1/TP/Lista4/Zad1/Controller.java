import javafx.fxml.FXML;
import javafx.scene.control.Button;
import javafx.scene.control.Label;
import javafx.scene.control.TextField;

public class Controller {
    @FXML
    private TextField input;
    @FXML
    private Label trojkatLabel;
    @FXML
    private Button policzButton;

    public void PoliczTrojkat() {
        int n = Integer.parseInt(input.getText());
        TrojkatPascala trojkat = new TrojkatPascala(n);
        trojkatLabel.setText(trojkat.printTriangle());
    }
}
