import React, { FC, useEffect } from 'react';
import { makeStyles } from '@material-ui/core/styles';
import { Content, Header, Page, pageTheme } from '@backstage/core';
import AccountCircleIcon from '@material-ui/icons/AccountCircle';
import SaveIcon from '@material-ui/icons/Save'; // icon save
import HealingIcon from '@material-ui/icons/Healing';
import SearchIcon from '@material-ui/icons/Search';
import EventIcon from '@material-ui/icons/Event';
import LocalHotelIcon from '@material-ui/icons/LocalHotel';
import EnhancedEncryptionIcon from '@material-ui/icons/EnhancedEncryption';
import ApartmentIcon from '@material-ui/icons/Apartment';
import Typography from '@material-ui/core/Typography';
import {
  Container,
  Grid,
  FormControl,
  Select,
  InputLabel,
  MenuItem,
  TextField,
  Button,
  Tab,
  Table,
} from '@material-ui/core';
import Link from '@material-ui/core/Link';
import { Link as RouterLink } from 'react-router-dom';
import { DefaultApi } from '../../api/apis'; // Api Gennerate From Command

const HeaderCustom = {
  minHeight: '50px',
};
// css style 
const useStyles = makeStyles(theme => ({
  root: {
    flexGrow: 1,
  },
  paper: {
    marginTop: theme.spacing(2),
    marginBottom: theme.spacing(2),
  },
  formControl: {
    width: 300,
  },
  selectEmpty: {
    marginTop: theme.spacing(2),
  },
  container: {
    display: 'flex',
    flexWrap: 'wrap',
  },
  textField: {
    width: 300,
  },
}));

const home: FC<{}> = () => {
    const classes = useStyles();
    const http = new DefaultApi();
  
  return (
    <Page theme={pageTheme.service}>
      <Header style={HeaderCustom} title={`SA63 HOSPITAL`}>
      <AccountCircleIcon aria-controls="fade-menu" aria-haspopup="true"  fontSize="large" />
        <div style={{ marginLeft: 10 }}> </div>
        <Link component={RouterLink} to="/">
             Logout
         </Link>
      </Header>
      <Content>
        <Container maxWidth="sm">

            <Grid item xs={1}></Grid>
          <Grid container spacing={3}>
            <Grid item xs={1}></Grid>
            
            <Grid item xs={1}>
              <div className={classes.paper}> </div>
            </Grid>
            <Grid item xs={9}>
   <Table >
     <tr>
       <td  align="center">
         <ApartmentIcon color="primary" style={{ fontSize: 100 }} />
     <Typography component="h1" variant="h5">
       SA63 HOSPITAL
     </Typography></td>
     </tr> &nbsp;&nbsp;
     <tr>
         <td><hr ></hr></td>
     </tr>
     </Table>
            </Grid>
            
            
            <Grid item xs={3}>
              <div className={classes.paper}> </div>
            </Grid>
            
             <Table>
             <tr>
         <td  align="center">
         <EnhancedEncryptionIcon color="primary" style={{ fontSize: 70 }} />
           <br></br><Typography component={RouterLink} to="/">
           Doctor'Register
         </Typography>
         </td>&nbsp;&nbsp;&nbsp;&nbsp;
         &nbsp;&nbsp;&nbsp;&nbsp;
         
         <td align="center">
         <LocalHotelIcon color="primary" style={{ fontSize: 70 }} />
         <br></br><Typography component={RouterLink} to="/">
           Patient'Register 
           </Typography>
         </td>&nbsp;&nbsp;&nbsp;&nbsp;
         &nbsp;&nbsp;&nbsp;&nbsp;
         <td align="center"> 
         <HealingIcon color="primary" style={{ fontSize: 70 }} />
         <br></br><Typography component={RouterLink} to="/WelcomePage">
           DrugAllergy
         </Typography></td>&nbsp;&nbsp;&nbsp;&nbsp;
         &nbsp;&nbsp;&nbsp;&nbsp;
         <td align="center">
         <SearchIcon color="primary" style={{ fontSize: 70 }} />
         <br></br><Typography component={RouterLink} to="/">
           DiagnoseSystem
         </Typography></td>&nbsp;&nbsp;&nbsp;&nbsp;
         &nbsp;&nbsp;&nbsp;&nbsp;
         <td align="center">
         <EventIcon color="primary" style={{ fontSize: 70 }} />
         <br></br><Typography component={RouterLink} to="/">
         AppointmentSystem
         </Typography></td>&nbsp;&nbsp;&nbsp;&nbsp;
         &nbsp;&nbsp;&nbsp;&nbsp;
       </tr>
   </Table>
            </Grid> 
        </Container>
      </Content>
    </Page>
  );
};
export default home;
