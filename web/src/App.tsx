import { useEffect } from 'react'

import './App.css'
import axios from 'axios'
import Logo from '@/components/Logo'
import { Counter } from '@/components/Counter'

function App() {
  

  useEffect(() => {
    axios.get('/api/users').then((response: any) => {
      console.log(response)
    })
  }, [])

  return (
    <>
      <div>
        <Logo />
      </div>
      <h1>Vite + React</h1>

      <Counter/>
    </>
  )
}

export default App
