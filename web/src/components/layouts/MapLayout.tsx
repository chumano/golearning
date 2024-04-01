import { Outlet } from "react-router-dom";


function MapLayout() {
  return (
    <div>MapLayout
        <div>
            <Outlet/>
        </div>
    </div>
  )
}

export default MapLayout