import Container from "react-bootstrap/Container";
import Nav from "react-bootstrap/Nav";
import Navbar from "react-bootstrap/Navbar";

export function NavBar() {
  return (
    <Navbar expand="lg" className="display-3">
      <Container>
        <Navbar.Brand>Purse</Navbar.Brand>
        <Navbar.Toggle />
        <Navbar.Collapse></Navbar.Collapse>
        <Nav>
          <Nav.Link href="#home">Home</Nav.Link>
          <Nav.Link href="#link">Link</Nav.Link>
        </Nav>
      </Container>
    </Navbar>
  );
}
