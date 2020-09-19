import React from 'react';
import './App.css';
import { BrowserRouter, Switch, Route, useParams } from 'react-router-dom';
import Home from "./component/home/home.js";
import Detail from "./component/detail/detail.js";

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
