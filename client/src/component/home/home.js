import React from 'react';
import corejs from "./assets/core.js";
import css from "./assets/asset.css";
import CssBaseline from '@material-ui/core/CssBaseline';
import Typography from '@material-ui/core/Typography';
import Container from '@material-ui/core/Container';
import { Link } from 'react-router-dom';
import Grid from '@material-ui/core/Grid';
import axios from "axios";

class Home extends React.Component{
  constructor(props){
    super(props)
     this.state = {
      url : "http://localhost:9000",
      data : null,
      count: null,
      next : null,
      previous: null,
      results : []
    }
  }
  componentDidMount(){
    axios.get(this.state.url+"/pokemon/get/9")
    .then((res) => {
          this.setState({data: res.data,count:res.data.count,next:res.data.next,previous:res.data.previous,results:res.data.results})
    })
  }

  cards = () => {
    return this.state.results.map(r => (
      <Grid item xs={12} md={4}>
      <Link to={"detail/" +  r.name} >
      <div class="card1 1">
        <div class="card_image"> <img src="https://assets.pokemon.com//assets/cms2/img/misc/virtual-backgrounds/sword-shield/galar-scenery.png" /> </div>
        <div class="card_title title-white">
          <p>{r.name}</p>
        </div>
      </div>
      </Link>
      </Grid>
    ))
  }

  render(){
    return(
      <div>
        <CssBaseline />
        <Grid container spacing={3}>
          {this.state.results.length == 0
            ? 'Loading results...': this.cards()
          }
        </Grid>
      </div>
    )
  }
}
export default Home;
