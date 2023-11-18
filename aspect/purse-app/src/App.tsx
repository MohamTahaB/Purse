import Button from "react-bootstrap/Button";
import "bootstrap/dist/css/bootstrap.min.css";
import { NavBar } from "./components/NavBar";

function App() {
  return (
    <>
      <NavBar/>
      <Button variant="primary">Primary</Button>{" "}
    </>
  );
}

export default App;
