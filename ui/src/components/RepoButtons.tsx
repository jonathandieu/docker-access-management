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


  const getRepositories = async () => {
    const result = await ddClient.extension.vm?.service?.get('/repositories?namespace=chefjon&max_results=25');
    setResponse(JSON.stringify(result));
    var obj = JSON.parse(JSON.stringify(result)); // JSON -> string -> JS Object
    var data = Object.values(obj)[0]; // JS Object -> Array -> JS Object
    setHeaders(Object.keys(data[0])); // table header
    setBodies(Object.values(data[0])); // table body
  };

  const getRepository = async () => {
    const result = await ddClient.extension.vm?.service?.get('/repository?namespace=chefjon&name=personal-repo-demo');
    setResponse(JSON.stringify(result));
  };

  const createRepository = async () => {
    const result = await ddClient.extension.vm?.service?.post('/repository?namespace=chefjon&name=test-this', "");
    setResponse(JSON.stringify(result));
  };

  const deleteRepository = async () => {
    const result = await ddClient.extension.vm?.service?.delete('/repository?namespace=chefjon&name=test-this');
    setResponse(JSON.stringify(result));
  };

    return (
        <>
     <div>
        <Typography variant="h2" component="div">Repository Management</Typography>
        <Stack
        direction="row"
        alignItems="start"
        spacing={2}
        sx={{ mt: 4 }}
         mb={2}
        divider={<Divider orientation="vertical" flexItem />}>
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

        </Stack>
        <TextField
        label=" Repo Status response"
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