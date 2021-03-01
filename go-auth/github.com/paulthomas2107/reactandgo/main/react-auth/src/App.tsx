import React from "react";
import "./App.css";
import { BrowserRouter, Route } from "react-router-dom";
import Login from "./pages/Login";
import Nav from "./navigation/Nav";
import Home from "./pages/Home";
import Register from "./pages/Register";

function App() {
  return (
    <div className="App">
      <BrowserRouter>
        <Nav />
        <main className="form-signin">
          <Route path="/" exact component={Home} />
          <Route path="/login" component={Login} />
          <Route path="/register" component={Register} />
        </main>
      </BrowserRouter>
    </div>
  );
}

export default App;
