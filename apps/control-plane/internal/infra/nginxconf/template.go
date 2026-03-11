package nginxconf

// TODO: MOVE TO DB
const serverTemplate = `
server {
    listen 80;
    server_name {{ .Domain }};

    location / {
        proxy_pass {{ .Upstream }};
        proxy_set_header Host $host;
        proxy_set_header X-Real-IP $remote_addr;
    }
}
`
