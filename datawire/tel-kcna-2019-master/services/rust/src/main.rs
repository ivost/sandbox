extern crate simple_server;

use simple_server::{Method, Server, StatusCode};

fn get_value() -> String {
    let mut response = match reqwest::get("http://base") {
        Ok(response) => response,
        Err(e) => return e.to_string(),
    };
    let value = match response.text() {
        Ok(value) => value,
        Err(e) => return e.to_string(),
    };
    return value;
}

fn get_hostname() -> String {
    let hostname = match gethostname::gethostname().into_string() {
        Ok(hostname) => hostname,
        Err(_) => "Invalid UTF-8".to_string(),
    };
    return hostname;
}

fn index() -> String {
    let value = "salty".to_string() + &get_value();
    let hashcode = base64::encode(&value);
    let lines = vec![
        "[ Hello KubeCon NA 2019! ]".to_string(),
        "[ Greetings from Rust    ]".to_string(),
        format!("[ Code: {} ]", hashcode),
        "".to_string(),
        format!("Host: {}", get_hostname()),
        format!("Now:  {}", chrono::Utc::now()),
    ];
    let result = lines.join("\n") + "\n";
    return result;
}

fn main() {
    let server = Server::new(|request, mut response| {
        println!("* {} {}", request.method(), request.uri());

        match (request.method(), request.uri().path()) {
            (&Method::GET, "/") => {
                let body = index();
                println!("  => OK");
                Ok(response.body(body.as_bytes().to_vec())?)
            }
            (_, _) => {
                println!("  => 404");
                response.status(StatusCode::NOT_FOUND);
                Ok(response.body("404 Not found".as_bytes().to_vec())?)
            }
        }
    });

    server.listen("0.0.0.0", "8000");
}
