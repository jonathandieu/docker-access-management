import { Box, Button, Drawer, List, ListItem, ListItemButton, ListItemIcon, ListItemText, Menu, MenuItem, Toolbar, Typography } from '@mui/material';
import React from 'react';

type Anchor = 'top';
export default function Navbar() {
   const [state, setState] = React.useState({
    top: false,
   });

   const toggleDrawer =
   (anchor: Anchor, open: boolean) =>
   (event: React.KeyboardEvent | React.MouseEvent) => {
     if (
       event.type === 'keydown' &&
       ((event as React.KeyboardEvent).key === 'Tab' ||
         (event as React.KeyboardEvent).key === 'Shift')
     ) {
       return;
     }

     setState({ ...state, [anchor]: open });
   };

   const list = (anchor: Anchor) => (
    <Box 
    sx={{width: anchor === 'top' ? 'auto' : 250}}
    role="demo"
    onClick={toggleDrawer(anchor, false)}
    onKeyDown={toggleDrawer(anchor, false)}
    >
        <List>
            {['Repository', 'Organization'].map((text, index) => (
                <ListItem key={text} disablePadding>
                  <ListItemButton>
                    <ListItemIcon>
                        {/* {index % 2 === 0 ? <PhotoCameraIcon /> : <CorporateFareIcon/>} */}
                    </ListItemIcon>
                    <ListItemText primary={text} />
                    </ListItemButton> 
                </ListItem>
            ))}
        </List>
    </Box>
   )
   
    return (
    <>
     <div>
      {(['top'] as const).map((anchor) => (
        <React.Fragment key={anchor}>
          <Button onClick={toggleDrawer(anchor, true)}>D.A.M</Button>
          <Drawer
            anchor={anchor}
            open={state[anchor]}
            onClose={toggleDrawer(anchor, false)}
          >
            {list(anchor)}
          </Drawer>
        </React.Fragment>
      ))}
    </div>
    </>
    )
}