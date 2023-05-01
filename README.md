# traefik-plugin-static-sites

Simple plugin for Traefik for serving static sites from MinIO buckets.

This plugin updates the path of a request with hostname from original request:

`https://a.example.com/ ->  https://<minio>/a.example.com/index.html`


````

services:
  traefik:
    image: traefik
    # other parameters are omitted for readability purposes 
    command:
      - "--experimental.plugins.traefik-plugin-static-sites.modulename=github.com/russ-p/traefik-plugin-static-sites"
      - "--experimental.plugins.traefik-plugin-static-sites.version=v0.1.5"    

  minio:
    image: minio/minio
    # other parameters are omitted for readability purposes 
    labels:
      - "traefik.enable=true"
      - "traefik.http.services.minio.loadbalancer.server.port=9000"
      - "traefik.http.middlewares.minio-site.plugin.traefik-plugin-static-sites.SpaFriendly=false"      
      - "traefik.http.routers.minio_static_site.rule=Host(`example.com`) || Host(`a.example.com`) || Host(`b.example.com`)"
      - "traefik.http.routers.minio_static_site.entrypoints=web,websecure"   
      - "traefik.http.routers.minio_static_site.middlewares=minio-site"  
      - "traefik.http.routers.minio_static_site.service=minio"
````