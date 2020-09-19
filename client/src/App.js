import React from 'react';
import './App.css';
import CssBaseline from '@material-ui/core/CssBaseline';
import Typography from '@material-ui/core/Typography';
import { BrowserRouter, Switch, Route, useParams } from 'react-router-dom';
import Container from '@material-ui/core/Container';
import Grid from '@material-ui/core/Grid';
import Home from "./component/home/home.js";
import Detail from "./component/detail/detail.js";
import axios from "axios";

class App extends React.Component{
  constructor(props){
    super(props)
     this.state = {
       id : null
    }
  }

  render(){
    console.log(this.props)
    return(
      <div>
        <BrowserRouter>
        <Switch>
           <Route path="/detail/:id">
             <Detail id={this.state.id} />
           </Route>
           <Route path="/">
             <Home />
           </Route>
         </Switch>
         </BrowserRouter>
      </div>
    )
  }
}
export default App;
