import { useState, useEffect } from "react";
import { Container, FormControl, Select, SimpleGrid } from "@chakra-ui/react";
import countryCode from "../database/country_code.json";
import { api, errorHandler } from "../services/Utils";
import * as MusixmatchArtist from "../types/Artist";
import { ArtistModal } from "../components/ArtistModal";
import User from "../types/User";
import { getUser, isAuth } from "../services/Auth";
import { useNavigate } from "react-router-dom";

type countryCodeType = {
  [key: string]: string;
};

const _countryCode: countryCodeType = countryCode;

export const HomeView: React.FC = () => {
  const [user, setUser] = useState<User>();
  const [countryCode, setCountryCode] = useState<string>("TW");
  const [artistList, setArtistList] = useState<MusixmatchArtist.Artistlist2[]>(
    []
  );
  const navigate = useNavigate();

  useEffect(() => {
  //   async function getPost() {
  //     try {
  //       const response = await api.get<MusixmatchArtist.HttpResponse>(
  //         `/musixmatch/getLastTenArtistsByCountryCode/${countryCode}`
  //       );

  //       const _result = response.data;
  //       setArtistList(_result.data.result.message.body.artist_list);
  //     } catch (error) {
  //       errorHandler(error as any, {});
  //     }
  //   }
  //   getPost();

        setUser(getUser());
  }, [countryCode]);

  if (!isAuth()) {
    navigate("/login");
    return <></>;
  }

  return (
    <Container>
      <h1>Hello {user?.name}</h1>
      {/* <FormControl>
        <Select
          placeholder="Select country"
          value={countryCode}
          onChange={(event) => setCountryCode(event.target.value)}
        >
          {Object.keys(_countryCode).map((code) => {
            return (
              <option key={code + "-" + _countryCode[code]} value={code}>
                [{code}] {_countryCode[code]}
              </option>
            );
          })}
        </Select>
      </FormControl>

      <SimpleGrid pt={8} pb={16} columns={1} spacing={10}>
        {artistList.map((a, k) => {
          return <ArtistModal artist={a.artist} level={k}></ArtistModal>;
        })}
      </SimpleGrid> */}
    </Container>
  );
};

export default HomeView;
