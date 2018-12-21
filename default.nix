{ stdenv, buildGoPackage }:

buildGoPackage rec {
  name = "dumpster";

  goPackagePath = "github.com/rvolosatovs/dumpster";

  src = ./.;

  meta = with stdenv.lib; {
    description = "Dump files";
    license = licenses.mit;
    homepage = https://github.com/rvolosatovs/dumpster;
  };
}
