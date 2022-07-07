import React from 'react';
import Button from '@mui/material/Button';
import { createDockerDesktopClient } from '@docker/extension-api-client';
import { Stack, Table, TableBody, TableHead, TableRow, TextField, Typography } from '@mui/material';
import Navbar from './components/Navbar';
import Footer from './components/Footer';
import { stringify } from 'querystring';
import Tester from './components/Tester';
// Note: This line relies on Docker Desktop's presence as a host application.
// If you're running this React app in a browser, it won't work properly.
const client = createDockerDesktopClient();

function useDockerDesktopClient() {
  return client;
}

export function App() {
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


  return (
    <>
    <Navbar />
      <Typography variant="h3">Docker Access Management (DAM)</Typography>
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
      <Tester />
      <Tester />
      <div>
      <Table>
          <TableRow>
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
