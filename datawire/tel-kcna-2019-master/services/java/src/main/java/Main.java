import static spark.Spark.get;
import static spark.Spark.port;

import java.io.BufferedReader;
import java.io.InputStreamReader;
import java.net.HttpURLConnection;
import java.net.InetAddress;
import java.net.URL;
import java.net.UnknownHostException;
import java.time.Instant;
import java.util.Base64;

public class Main {
    private static String getValue() {
        try {
            URL url = new URL("http://base");
            HttpURLConnection conn = (HttpURLConnection) url.openConnection();
            conn.setRequestMethod("GET");
            BufferedReader in = new BufferedReader(new InputStreamReader(conn.getInputStream()));
            String inputLine;
            StringBuilder content = new StringBuilder();
            while ((inputLine = in.readLine()) != null) {
                content.append(inputLine);
            }
            in.close();
            return content.toString();
        } catch (Exception e) {
            return e.toString();
        }
    }

    private static String getHostname() {
        InetAddress ip;
        try {
            ip = InetAddress.getLocalHost();
            return ip.getHostName();
        } catch (UnknownHostException e) {
            return e.toString();
        }
    }

    public static void main(String[] args) {
        port(8000);
        get("/", (req, response) -> {
            // Get information from other services
            String value = "salty" + getValue();

            // Perform some computation
            String hashcode = Base64.getEncoder().encodeToString(value.getBytes());
            String[] lines = {
                    "[ Hello KubeCon NA 2019! ]",
                    "[ Greetings from Java    ]",
                    String.format("[ Code: %s ]", hashcode),
                    "",
                    String.format("Host: %s", getHostname()),
                    String.format("Now:  %s", Instant.now()),
            };
            String res = String.join("\n", lines) + "\n";

            // Return the result
            response.type("text/plain");
            return res;
        });
    }
}
