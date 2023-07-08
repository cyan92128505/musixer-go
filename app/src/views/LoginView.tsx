import {  useState } from "react";
import { useNavigate } from "react-router";
import { Link as ReachLink } from "react-router-dom";

import { api, errorHandler } from "../services/Utils";
import { isAuth, login } from "../services/Auth";
import {
  Grid,
  VStack,
  Input,
  Box,
  Button,
  Container,
  Stack,
  Link,
} from "@chakra-ui/react";
import { ColorModeSwitcher } from "../ColorModeSwitcher";
import { Logo } from "../Logo";
import { useForm } from "react-hook-form";
import AlertPop from "../components/registerForm/AlertPop";

type LoginFormType = {
  email: string;
  password: string;
};


export const LoginView: React.FC = () => {
    const {
      register,
      handleSubmit,
      formState: { errors },
    } = useForm<LoginFormType>();
    const [data, setData] = useState<LoginFormType>();

  const navigate = useNavigate();
  const onSubmit = async (data: LoginFormType) => {
    await api
      .post("/auth/signin", data)
      .then((response) => login(response))
      .catch((error) => {
        errorHandler(error, {});
      });

    setData(data);
  };


  // When the user already authenticated
  if (isAuth()) {
    navigate("/");
    return <></>;
  }

  return (
    <>
      <form onSubmit={handleSubmit(onSubmit)}>
        <Box textAlign="center" fontSize="xl">
          <Grid minH="100vh" p={3}>
            <ColorModeSwitcher justifySelf="flex-end" />
            <VStack spacing={8}>
              <Container maxW="md">
                <Stack padding="4" spacing={3}>
                  <Logo></Logo>
                  <Input
                    type="email"
                    placeholder="Last name"
                    {...register("email", {
                      required: "Please enter email",
                      minLength: 3,
                      maxLength: 100,
                    })}
                  />{" "}
                  {errors.email && <AlertPop title={errors.email.message} />}
                  <Input
                    type="password"
                    placeholder="Password"
                    {...register("password", {
                      required: "Please enter Password",
                      minLength: { value: 8, message: "Too short" },
                    })}
                  />{" "}
                  {errors.password && (
                    <AlertPop title={errors.password.message} />
                  )}
                  <br />
                  <Button colorScheme="teal" type="submit" fontSize="19px">
                    Login
                  </Button>
                  <Link as={ReachLink} to="/register">
                    Register
                  </Link>
                </Stack>
              </Container>
            </VStack>
          </Grid>
        </Box>
      </form>
    </>
  );
};

export default LoginView;
