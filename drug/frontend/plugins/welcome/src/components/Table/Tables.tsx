import React, { useState, useEffect } from 'react';
import { makeStyles } from '@material-ui/core/styles';
import Table from '@material-ui/core/Table';
import TableBody from '@material-ui/core/TableBody';
import TableCell from '@material-ui/core/TableCell';
import TableContainer from '@material-ui/core/TableContainer';
import TableHead from '@material-ui/core/TableHead';
import TableRow from '@material-ui/core/TableRow';
import Paper from '@material-ui/core/Paper';
import Button from '@material-ui/core/Button';
import { DefaultApi } from '../../api/apis';
import { EntDrugAllergy } from '../../api/models/EntDrugAllergy';
import { Header, Page, pageTheme } from '@backstage/core';
import AccountCircleIcon from '@material-ui/icons/AccountCircle';
import {Container,Link, Grid} from '@material-ui/core';
import { Link as RouterLink } from 'react-router-dom';
const useStyles = makeStyles({
table: {
minWidth: 650,
},
});
const HeaderCustom = {
    minHeight: '50px',
  };
export default function ComponentsTable() {
const classes = useStyles();
const api = new DefaultApi();

const [drugAllergys, setDrugAllergys] = React.useState<EntDrugAllergy[]>([]);
const [loading, setLoading] = useState(true);

useEffect(() => {
  const getDrugAllergys = async () => {
      const res = await api.listDrugAllergy({ limit: 20, offset: 0 });
      setLoading(false);
      setDrugAllergys(res);
      console.log(drugAllergys);
      console.log(res);
  };
  getDrugAllergys();
}, [loading]);

const deleteDrugAllergys = async (id: number) => {
const res = await api.deleteDrugAllergy({ id: id });
setLoading(true);
};

return (
    <Page theme={pageTheme.service}>
      <Header style={HeaderCustom} title={`Drug Allergy`}>
      <AccountCircleIcon aria-controls="fade-menu" aria-haspopup="true"  fontSize="large" />
        <div style={{ marginLeft: 10 }}> </div>
        <Link component={RouterLink} to="/">
             Logout
         </Link>
      </Header>
      <Grid item xs={12}></Grid>
<TableContainer component={Paper}>
<Table className={classes.table} aria-label="simple table">
<TableHead>
<TableRow>
<TableCell align="center">No.</TableCell>
<TableCell align="center">Doctor</TableCell>
<TableCell align="center">Patient</TableCell>
<TableCell align="center">Medicine</TableCell>
<TableCell align="center">Manner</TableCell>
<TableCell align="center"><Link component={RouterLink} to="/WelcomePage">
           <Button variant="contained" color="primary">
             Add Drug Allergy
           </Button>
         </Link></TableCell>
</TableRow>
</TableHead>

<TableBody>
{ drugAllergys.map((item:any) => (
<TableRow key={item.id}>
<TableCell align="center">{item.id}</TableCell>
<TableCell align="center">{item.edges?.doctor?.doctorName}</TableCell>
<TableCell align="center">{item.edges?.patient?.patientName}</TableCell>
<TableCell align="center">{item.edges?.medicine?.medicineName}</TableCell>
<TableCell align="center">{item.edges?.manner?.mannerName}</TableCell>
<TableCell align="center">
<Button
onClick={() => {
if (item.id === undefined){
return;
}
deleteDrugAllergys(item.id);
}}
style={{ marginLeft: 10 }}
variant="contained"
color="secondary"
>
Delete
</Button>
</TableCell>
</TableRow>
))}
</TableBody>
</Table>
</TableContainer>
</Page>
);
}