# servemydata (Spectral data server)

## Purpose

servemydata is a little playground for me to take another journey into the between software development and science. The
goal is to provide a simple REST API to serve spectral data from some persistence layer. This can be very helpful for
working groups having a shared set of experimental data that needs to be stored in a central place and retrieved,
filtered by certain criteria and distributed for usage in analytical tools, e.g. plotting.

### Design

The application is designed following (at least) roughly the approach of hexagonal architecture: a separation of the
core business logic from the technical environment the application runs in. The chose for Go as the programming language
for this project was made because it is very suitable to implement performant server-side web apps with only few lines
of code and for educational purposes as I need to improve my skills in Go.

## Next steps

- implement the core use cases for the first runnable version of the application
- write tests for the core functionality
- add the basic functionality using of an in-memory database
