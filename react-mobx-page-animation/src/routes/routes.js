import Home from "../components/views/Home";
import AnimationType from "../components/views/AnimationType";
import CommandList from "../components/views/CommandList";

const routes = [
  { path: "/home", exact: true, name: "Home", component: Home },
  {
    path: "/animation/",
    exact: true,
    name: "animation ",
    component: AnimationType
  },
  {
    path: "/animation/:animationType",
    name: "animation Type",
    component: AnimationType
  },
  {
    path: "/cl",
    name: "command ist",
    component: CommandList
  }
];

export default routes;
