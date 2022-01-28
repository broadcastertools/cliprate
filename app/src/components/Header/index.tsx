import React from "react";
import { AppBar, Toolbar } from '@mui/material';
import logo from './logo.png';

export const Header = () => {
    return <AppBar>
      <Toolbar sx={{py: 1}}>
        <img src={logo} alt="MuTeX's logo" />
      </Toolbar>
    </AppBar>
};
