import javafx.fxml.FXML;
import javafx.scene.control.Button;
import javafx.scene.control.Label;
import javafx.scene.control.TextField;
import java.io.*;


public class Controller {
    @FXML
    private Button policzButton;
    @FXML
    private Label trojkatLabel;
    @FXML
    private TextField input;
    @FXML
    public void policzTrojkat() throws IOException {
                    trojkatLabel.setText("");
                    String command = "./src/PascalT " + input.getText();
                    System.out.println(command);
                    Process process = Runtime.getRuntime().exec(command);
                    BufferedReader reader = new BufferedReader(new InputStreamReader(process.getInputStream()));
                    for(int i = 0; i < Integer.parseInt(input.getText()); i++){
                        trojkatLabel.setText(trojkatLabel.getText() + reader.readLine() + '\n');
                    }
        }
    }

