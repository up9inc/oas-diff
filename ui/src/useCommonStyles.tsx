import { Button, styled } from '@mui/material';
import { makeStyles } from '@mui/styles';

// @ts-ignore
export const useCommonStyles = makeStyles(() => ({
    select: {
        margin: "0px !important",
        width: "250px"
    }
}));

export const CollapseButton = styled(Button)({
    backgroundColor: "#205cf5",
    color: "white",
    fontWeight: "600 !important",
    fontSize: 12,
    padding: "7.5px 10px",
    borderRadius: "6px ! important",

    "&:hover": {
        backgroundColor: "#205cf5",
        cursor: "pointer"
    },
    "&:disabled": {
        backgroundColor: "rgba(0, 0, 0, 0.26)"
    }
});
