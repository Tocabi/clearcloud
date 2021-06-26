FROM scratch

EXPOSE 5555

ENV POSTGRES_HOST=localhost
ENV POSTGRES_DB=clearcloud
ENV POSTGRES_USER=clearcloud
ENV POSTGRES_PASSWORD=clearcloud
ENV POSTGRES_PORT=5432

ADD clearcloud /app/clearcloud

ENTRYPOINT ["/app/clearcloud"]