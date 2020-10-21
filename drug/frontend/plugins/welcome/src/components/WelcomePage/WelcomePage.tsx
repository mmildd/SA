import React, { FC, useEffect } from 'react';
import { makeStyles } from '@material-ui/core/styles';
import { Content, Header, Page, pageTheme } from '@backstage/core';
import AccountCircleIcon from '@material-ui/icons/AccountCircle';
import Swal from 'sweetalert2'
import SaveIcon from '@material-ui/icons/Save'; // icon save
import {
  Container,
  Grid,
  FormControl,
  Select,
  InputLabel,
  MenuItem,
  TextField,
  Button,
} from '@material-ui/core';
import Link from '@material-ui/core/Link';
import { Link as RouterLink } from 'react-router-dom';
import { DefaultApi } from '../../api/apis'; // Api Gennerate From Command
import { EntDoctor } from '../../api/models/EntDoctor'; // import interface Doctor
import { EntPatient } from '../../api/models/EntPatient'; // import interface Video
import { EntMedicine} from '../../api/models/EntMedicine'; // import interface Resolution
import { EntManner} from '../../api/models/EntManner'; // import interface Resolution
import { EntDrugAllergy, mapValues } from '../../api';


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
interface drugAllergy {
  manner: number;
  medicine: number;
  patient: number;
  doctor: number;
}

interface Manner {
  manner_name: string;
}

const DrugAllergy: FC<{}> = () => {
  const classes = useStyles();
  const http = new DefaultApi();
  const [drugAllergy, setDrugAllergy] = React.useState<
  Partial<drugAllergy>
>({});

const [doctors, setDoctors] = React.useState<EntDoctor[]>([]);
const [patients, setPatients] = React.useState<EntPatient[]>([]);
const [medicines, setMedicines] = React.useState<EntMedicine[]>([]);
const [manners, setManners] = React.useState<Partial<Manner>>({});



  const getDoctor = async () => {
    const res = await http.listDoctor({ limit: 2, offset: 0 });
    setDoctors(res); 
  };
  const getPatient = async () => {
    const res = await http.listPatient({ limit: 4, offset: 0 });
    setPatients(res);
  };

  const getMedicine = async () => {
    const res = await http.listMedicine ({ limit: 4, offset: 0 });
    setMedicines (res);
  };

  

  // Lifecycle Hooks
    useEffect(() => {
    getDoctor();
    getPatient();
    getMedicine();
    
   
  }, []);

  // set data to object DrugAllergy
  const handleChange = (
    event: React.ChangeEvent<{ name?: string; value: unknown }>,
  ) => {
    const name = event.target.name as keyof typeof DrugAllergy;
    const { value } = event.target;
    setDrugAllergy({ ...drugAllergy, [name]: value });
    
  };

  const handleChangemanner = (
    event: React.ChangeEvent<{ name?: string; value: unknown }>,
  ) => {
    const name = event.target.name as keyof typeof DrugAllergy;
    const { value } = event.target;
    setManners({ ...manners, [name]: value });
 
    
  };

  // clear input form
  function clear() {
    setDrugAllergy({});
  }

  // alert setting
  const Toast = Swal.mixin({
    toast: true,
    position: 'top-end',
    showConfirmButton: false,
    timer: 3000,
    timerProgressBar: true,
    didOpen: toast => {
      toast.addEventListener('mouseenter', Swal.stopTimer);
      toast.addEventListener('mouseleave', Swal.resumeTimer);
    },
  });

 // function save data
 function save() {
  
  const apiUrl2 = 'http://localhost:8080/api/v1/manners';
  const requestOptions2 = {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify(manners),
  };
  
  fetch(apiUrl2, requestOptions2)
    .then(response => response.json())
    .then(data => {
      
      drugAllergy.manner = data.id
      console.log(drugAllergy);
      http.createDrugAllergy({drugallergy:drugAllergy});
      console.log(data);
      if (data.status === false) {
      
        Toast.fire({
          icon: 'error',
          title: 'บันทึกข้อมูลไม่สำเร็จ',
        });
      } else {
        Toast.fire({
          icon: 'success',
          title: 'บันทึกข้อมูลสำเร็จ',
        });
      }
    });

   
}

  
  return (
    <Page theme={pageTheme.service}>
      <Header style={HeaderCustom} title={`Drug Allergy`}>
      <AccountCircleIcon aria-controls="fade-menu" aria-haspopup="true"  fontSize="large" />
        <div style={{ marginLeft: 10 }}> </div>
        <Link component={RouterLink} to="/">
             Logout
         </Link>
      </Header>
      <Content>
        <Container maxWidth="sm">

            <Grid item xs={3}></Grid>
          <Grid container spacing={3}>
            <Grid item xs={12}></Grid>
            
            <Grid item xs={3}>
              <div className={classes.paper}>Doctor </div>
            </Grid>
            <Grid item xs={9}>
              <FormControl variant="outlined" className={classes.formControl}>
              <InputLabel></InputLabel>
                <Select
                  name="doctor"
                  value={drugAllergy.doctor } // (undefined || '') = ''
                  onChange={handleChange}
                >
                  {doctors.map(item => {
                    return (
                      <MenuItem key={item.id} value={item.id}>
                        {item.doctorName}
                      </MenuItem>
                    );
                  })}
                </Select>
              </FormControl>
            </Grid>
            
            
            <Grid item xs={3}>
              <div className={classes.paper}>Patient</div>
            </Grid>
            <Grid item xs={9}>
              <FormControl variant="outlined" className={classes.formControl}>
                <InputLabel></InputLabel>
                <Select
                  name="patient"
                  value={drugAllergy.patient } // (undefined || '') = ''
                  onChange={handleChange}
                >
                  {patients.map(item => {
                    return (
                      <MenuItem key={item.id} value={item.id}>
                        {item.patientName}
                      </MenuItem>
                    );
                  })}
                </Select>
              </FormControl>
            </Grid>

            
            <Grid item xs={3}>
              <div className={classes.paper}>Medicine</div>
            </Grid>
            <Grid item xs={9}>
              <FormControl variant="outlined" className={classes.formControl}>
                <InputLabel></InputLabel>
                <Select
                  name="medicine"
                  value={drugAllergy.medicine } // (undefined || '') = ''
                  onChange={handleChange}
                >
                  {medicines.map(item => {
                    return (
                      <MenuItem key={item.id} value={item.id}>
                        {item.medicineName}
                      </MenuItem>
                    );
                  })}
                </Select>
              </FormControl>
            </Grid>

            
            <Grid item xs={3}>
              <div className={classes.paper}>Manner</div>
            </Grid>
            <Grid item xs={9}>
              <form className={classes.container} noValidate>
                <TextField
                  label=""
                  name="manner_name"
                  type="string"
                  multiline
                  variant="outlined"
                  rows={7}
                  style={{ width: 300 }}
                  value={manners.manner_name|| ''} // (undefined || '') = ''
                  className={classes.textField}
                  InputLabelProps={{
                    shrink: true,
                  }}
                  onChange={handleChangemanner}
                />
              </form>
            </Grid>

            
            <Grid item xs={3}></Grid>
            <Grid item xs={9}>
              <Button
                variant="contained"
                color="primary"
                size="large"
                startIcon={<SaveIcon />}
                onClick={save}
                component={RouterLink} to="/Tables"
              >
                Save
              </Button>
            </Grid>
          </Grid>

        </Container>
      </Content>
    </Page>
  );
};
export default DrugAllergy;
