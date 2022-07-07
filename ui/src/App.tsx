import React from 'react';
import Button from '@mui/material/Button';
import { createDockerDesktopClient } from '@docker/extension-api-client';
import { Stack, Table, TableBody, TableHead, TableRow, TextField, Typography } from '@mui/material';
import Navbar from './components/Navbar';
import { stringify } from 'querystring';
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
    var s = JSON.parse(JSON.stringify(result));
    var data = Object.values(s)[0];
    setHeaders(Object.keys(data[0]));
    setBodies(Object.values(data[0]));
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
      <Typography variant="h3">Docker extension demo</Typography>
      <Typography variant="body1" color="text.secondary" sx={{ mt: 2 }}>
        This is a basic page rendered with MUI, using Docker's theme. Read the
        MUI documentation to learn more. Using MUI in a conventional way and
        avoiding custom styling will help make sure your extension continues to
        look great as Docker's theme evolves.
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
      <Table>
          <TableRow>
            <TableHead>{ headers[0] }</TableHead>
            <TableHead>foo</TableHead>
          </TableRow>
          <TableRow>
            <TableBody>{ bodies }</TableBody>
            <TableBody>bar</TableBody>
          </TableRow>
        </Table>
      </div>
    </>
  );
}
