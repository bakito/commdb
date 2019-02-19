import 'bootstrap/dist/css/bootstrap.min.css';
import React, { Component } from "react";
import { render } from "react-dom";
import FormContainer from "./containers/FormContainer";


class App extends Component {
  render() {
    return (
      <div className="col-md-6">
        <h3> Sample Form Container </h3>
        <FormContainer />
      </div>
    );
  }
}

render(<App />, document.getElementById("root"));
