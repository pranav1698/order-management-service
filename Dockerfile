FROM centos
COPY order-management-service .
EXPOSE 8080
CMD [ "./order-management-service"]