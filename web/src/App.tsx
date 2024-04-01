import { lazy, useEffect } from 'react'

import './App.css'
import axios from 'axios'
import {
  createBrowserRouter,
  RouterProvider,
} from "react-router-dom";

import MainLayout from '@/components/layouts/MainLayout';
import { Button } from '@/components/ui/button';

const SettingsPage = lazy(() =>import('@/components/Counter'))

const router = createBrowserRouter([
  {
    path: "/",
    element: <MainLayout/>,
    children:[
      {
        path: "/",
        element: <div>
           <Button variant={'default'}>Click me</Button>
        </div>
      },
      {
        path: "/counter",
        element: <SettingsPage/>,
      }
    ]
  }
]);

function App() {
  
  useEffect(() => {
    axios.get('/api/users').then((response: any) => {
      console.log(response)
    })
  }, [])

  return (
    <RouterProvider router={router} />
  )
}

export default App
