import React, { Component, Fragment } from "react";
import { withRouter } from "react-router-dom";
import { inject, observer } from "mobx-react";

@inject("CommandModel")
@observer
class CommandList extends Component {
  _isMounted = false;
  constructor(props) {
    super(props);
  }

  componentWillUnmount() {
    this._isMounted = false;
  }

  componentDidMount() {
    this.props.CommandModel.load(this.props.CommandModel);
  }

  render() {
    const { rows } = this.props.CommandModel;
    return (
      <Fragment>
        <ul>
          {rows.map((r, idx) => {
            return <li>{idx + 1 + ". " + r.command}</li>;
          })}
        </ul>
      </Fragment>
    );
  }
}

export default withRouter(CommandList);
