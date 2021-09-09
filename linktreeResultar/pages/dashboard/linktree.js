import AdminMenu from "../../components/AdminMenu/AdminMenu";
import api from "../../services/api";
import LinkTreeDashboard from "../../components/LinkTreeDashboard";

function Links(props) {
  return (
    <div className={"dashboard"}>
      <AdminMenu file_path={props.file_path} />
      <LinkTreeDashboard />
    </div>
  );
}

Links.getInitialProps = async () => {
  let file_path = undefined;
  file_path = await api.get("read-logo").then((r) => {
    return r.data.file_path;
  }).catch((e) => {
    if (e.response.status == 400) return undefined;
  });
  return { file_path: file_path };
};

export default Links;
