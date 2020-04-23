/*!
 * @2018 Stanislav Shabalin
 */

import * as React from "react";
import { Rect } from "../model/Rect";
import TextField, {HelperText, Input} from '@material/react-text-field';
import MaterialIcon from '@material/react-material-icon';
import '@material/react-text-field/dist/text-field.css';

export interface AboutProps {screenRect: Rect}
export interface AboutState {}

export class About extends React.Component<AboutProps, AboutState> {
  constructor(props: AboutProps) {
    super(props)
  }
  render() {
    return (
      <div>
        <div>
          Screen size:
        </div>
        <div>
          width: {this.props.screenRect.getX2()}
        </div>
        <div>
          height: {this.props.screenRect.getY2()}
        </div>
        <TextField
          label='Dog'
          helperText={<HelperText>Help Me!</HelperText>}
          onTrailingIconSelect={() => this.setState({value: ''})}
          trailingIcon={<MaterialIcon role="button" icon="delete"/>}
        ><Input
           value=""
          />
        </TextField>
      </div>
    );
  }
}
