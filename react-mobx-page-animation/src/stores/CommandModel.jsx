import { action, autorun, observable } from "mobx";
import axios from "axios";

class CommandModel {
  @observable rows = [];

  @action
  load = m => {
    axios
      .get("/api/command")
      .then(function(response) {
        // handle success
        m.rows = response.data;
        console.log(response);
      })
      .catch(function(error) {
        // handle error
        console.log(error);
      });
  };
}

var store = (window.commandModel = new CommandModel());
export default new CommandModel();

autorun(() => {
  console.log("CommandModel.load: " + store.load);
});
