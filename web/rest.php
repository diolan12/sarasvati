<?php
include_once "config.php";
include_once "repo.php";

header('Content-Type: application/json; charset=utf-8');

$id = "";
$data = [];
if (!isset($_GET["id"])) {
    $data = getAllProvinces();
} else {
    $id = $_GET["id"];
    switch (strlen($id)) {
        case 2:
            // provinsi
            $data = getProvince($id);
            break;

        case 4:
            // kabupaten/kota
            $data = getRegency($id);
            break;

        case 7:
            // kecamatan substr($regencyID, 2);
            $data = getDistrict($id);
            break;
    }
}
echo json_encode($data);