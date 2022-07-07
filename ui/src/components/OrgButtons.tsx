import {  Button, Divider, Stack, TextField, Typography } from '@mui/material';
import React from 'react';
import { createDockerDesktopClient } from '@docker/extension-api-client';

const client = createDockerDesktopClient();

function useDockerDesktopClient() {
  return client;
}

export default function Tester() {
  const [response, setResponse] = React.useState<string>();
  const ddClient = useDockerDesktopClient();
  const [headers, setHeaders] = React.useState<string[]>();
  const [bodies, setBodies] = React.useState<string[]>();


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
     <div>
        <Typography variant="h2" component="div">Organization Management</Typography>
        <Stack
        direction="row"
        alignItems="start"
        spacing={2}
        sx={{ mt: 4 }}
        divider={<Divider orientation="vertical" flexItem />}>

         <Button variant="contained" onClick={getOrganizations}>
          Get Organizations
        </Button>

        <Button variant="contained" onClick={getOrganization}>
          Get Organization
        </Button>

        <Button variant="contained" onClick={createOrganizations}>
          Create Organization
        </Button>



        </Stack>
        <br></br>
        <TextField
        label=" Repo Status Response"
        sx={{ width: 670 }}
        disabled
        multiline
        variant="outlined"
        minRows={5}
        value={response ?? ''}
        />
    </div>
    </>
    )
}