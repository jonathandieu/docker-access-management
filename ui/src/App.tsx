import React from 'react';
import Button from '@mui/material/Button';
import { createDockerDesktopClient } from '@docker/extension-api-client';
import { Card, CardContent, Paper, Stack, Table, TableBody, TableCell, TableContainer, TableHead, TableRow, TextField, Typography } from '@mui/material';
import Navbar from './components/Navbar';
import { stringify } from 'querystring';
import ShieldIcon  from '@mui/material/Icon';
// Note: This line relies on Docker Desktop's presence as a host application.
// If you're running this React app in a browser, it won't work properly.
const client = createDockerDesktopClient();

function useDockerDesktopClient() {
  return client;
}

export function App() {
  const [response, setResponse] = React.useState<string>();
  const ddClient = useDockerDesktopClient();
  const [headers, setHeaders] = React.useState<string[]>();
  const [bodies, setBodies] = React.useState<string[]>();


  // 	// Repos routes
	// router.GET("/repositories", c.GetRepositories)
	// router.GET("/repository", c.GetRepository)
	// router.POST("/repository", c.CreateRepository)
	// router.DELETE("/repository", c.DeleteRepository)

	// // Orgs routes
	// router.GET("/organizations", c.GetOrganizations)
	// router.GET("/organization", c.GetOrganization)
	// router.GET("/organization", c.GetOrganization)

  const getRepositories = async () => {
    const result = await ddClient.extension.vm?.service?.get('/repositories?namespace=ryanhristovski&max_results=25');
    setResponse(JSON.stringify(result));
    var obj = JSON.parse(JSON.stringify(result)); // JSON -> string -> JS Object
    var data = Object.values(obj)[0]; // JS Object -> Array -> JS Object
    setHeaders(Object.keys(data[0])); // table header
    setBodies(Object.values(data[0])); // table body
  };

  const getRepository = async () => {
    const result = await ddClient.extension.vm?.service?.get('/repository?namespace=ryanhristovski&name=personal-repo-demo');
    setResponse(JSON.stringify(result));
  };

  const createRepository = async () => {
    const result = await ddClient.extension.vm?.service?.post('/repository?namespace=ryanhristovski&name=test-this', "");
    setResponse(JSON.stringify(result));
  };

  const deleteRepository = async () => {
    const result = await ddClient.extension.vm?.service?.delete('/repository?namespace=ryanhristovski&name=test-this');
    setResponse(JSON.stringify(result));
  };

  const getOrganization = async () => {
    const result = await ddClient.extension.vm?.service?.get('/organization?org_name=dockerhackathon');
    setResponse(JSON.stringify(result));
  };

  const getOrganizations = async () => {
    const result = await ddClient.extension.vm?.service?.get('/organizations?username=ryanhristovski&max_results=25');
    setResponse(JSON.stringify(result));
  };

  const createOrganizations = async () => {
    const result = await ddClient.extension.vm?.service?.post('/organization?org_name=test-create&company=dam', "");
    setResponse(JSON.stringify(result));
  };


  return (
    <>
    <Navbar />
      <Typography variant="h3">Docker Access Management D.A.M</Typography>
      <Typography variant="body1" color="text.secondary" sx={{ mt: 2 }}>
      Is the process of having to open a new browser, navigating to Dockerhub.com, and signing in, 
      all just to access your repositories and organizations a pain?
      Have you ever wished that this could all be done from Docker Desktop? Wish no longer!
      </Typography>
      <Typography variant="body1" color="text.secondary" sx={{ mt: 2 }}>
        Pressing the below button will trigger a request to the backend. Its
        response will appear in the textarea.
      </Typography>

      <Stack direction="row" alignItems="start" spacing={2} sx={{ mt: 4 }}>
        <Button variant="contained" onClick={getRepository}>
          Get Repository
        </Button>

        <Button variant="contained" onClick={getRepositories}>
          Get Repositories
        </Button>      

        <Button variant="contained" onClick={createRepository}>
          Create Repository
        </Button>         

        <Button variant="contained" onClick={deleteRepository}>
          Delete Repository
        </Button>         

        <Button variant="contained" onClick={getOrganizations}>
          Get Organizations
        </Button>

        <Button variant="contained" onClick={getOrganization}>
          Get Organization
        </Button>        

        <Button variant="contained" onClick={createOrganizations}>
          Create Organization
        </Button>        

        <TextField
          label="Orgs response"
          sx={{ width: 480 }}
          disabled
          multiline
          variant="outlined"
          minRows={5}
          value={response ?? ''}
        />       

      </Stack>

      <div>
        <Typography variant="h2">Get Repos Data:</Typography>
        <TableContainer component={Paper}>
      <Table>
          <TableRow>
            <TableCell>{headers[0] } : {bodies} </TableCell>
          </TableRow>
        </Table>
        </TableContainer>
      </div>

      <div>
        <Typography sx={{marginTop: 5}}variant="h2"> Example: </Typography>
      <TableContainer component={Paper}>
        <Table sx={{minWidth: 650 }} aria-label='repo-table'>
          <TableHead>
            <TableRow sx={{marginBottom:1}}>
              <TableCell> mobythewhale/my-first-repo </TableCell>
            </TableRow>
            <TableRow>
           <TableCell>mobythewhale/my-second-repo</TableCell>   
            </TableRow>
            <TableRow>
              <TableCell>mobythewhale/my-third-repo</TableCell>
            </TableRow>
          </TableHead>
          <TableBody>
          </TableBody>
        </Table>
      </TableContainer>
      </div>

    </>
  );
}
