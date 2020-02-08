const os = require("os");

const got = require('got');
const express = require("express");

const app = express();

async function getValue() {
    try {
        const response = await got("http://base");
        return response.body;
    } catch (err) {
        return err.toString();
    }
}

app.get("/", async (_, response) => {
    const value = "salty" + await getValue();
    const hashvalue = Buffer.from(value).toString("base64");
    const lines = [
        "[ Hello KubeCon NA 2019! ]",
        "[ Greetings from NodeJS  ]",
        `[ Code: ${hashvalue} ]`,
        "",
        `Host: ${os.hostname()}`,
        `Now:  ${(new Date()).toISOString()}`,
    ];
    const res = lines.join("\n") + "\n";
    response.type("text/plain");
    response.send(res);
});

app.listen(8000);
