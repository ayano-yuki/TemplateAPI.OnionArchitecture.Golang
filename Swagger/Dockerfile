FROM swaggerapi/swagger-ui

COPY ./docs/swagger.json /usr/share/nginx/html/swagger.json

ENV SWAGGER_JSON=/usr/share/nginx/html/swagger.json

EXPOSE 8080
CMD ["nginx", "-g", "daemon off;"]
