<?php
include_once "config.php";

function getAllProvinces()
{
    return  [
        'context' => null,
        'child' => loadIndex(RootDir, INDEX_MAIN),
    ];
}



// 35
function getProvince(string $provinceID)
{
    $data = [
        'context' => null,
        'child' => []
    ];
    $provinces = getAllProvinces()['child'];
    foreach ($provinces as $province) {
        if ($provinceID === $province['id']) {
            $provincePath = RootDir.DIRECTORY_SEPARATOR.$province['id'].'-'.$province['name'];
            $data['context'] = $province;
            $data['child'] = loadIndex($provincePath, INDEX_MAIN);
            return $data;
        }
    }
    return null;
}

// 3510
function getRegency(string $regencyID)
{
    $data = [
        'context' => null,
        'child' => []
    ];
    $provinceID = substr($regencyID, 0, 2);
    $province = getProvince($provinceID);
    foreach ($province['child'] as $regency) {
        if ($regencyID === $regency['id']) {
            $provincePath = RootDir.DIRECTORY_SEPARATOR.$province['context']['id'].'-'.$province['context']['name'];
            $regencyPath = $provincePath.DIRECTORY_SEPARATOR.$regency['id'].'-'.$regency['name'];
            $data['context'] = $regency;
            $data['child'] = loadIndex($regencyPath, INDEX_MAIN);
            $data['ponpes'] = loadIndex($regencyPath, INDEX_PONPES);
            // $data['university'] = loadIndex($regencyPath, INDEX_UNIVERSITY);
            return $data;
        }
    }
    return null;
}

// 3510070
function getDistrict($districtID) {
    $data = [
        'context' => null,
        'child' => [],
        'kb' => [],
        'tk' => [],
        'sd' => [],
        'smp' => [],
        'sma' => [],
        'smk' => [],
        'slb' => [],
    ];
    $provinceID = substr($districtID, 0, 2);
    $province = getProvince($provinceID);
    $regencyID = substr($districtID, 0, 4);
    $regency = getRegency($regencyID);
    foreach ($regency['child'] as $district) {
        if ($districtID === $district['id']) {
            $provincePath = RootDir.DIRECTORY_SEPARATOR.$province['context']['id'].'-'.$province['context']['name'];
            $regencyPath = $provincePath.DIRECTORY_SEPARATOR.$regency['context']['id'].'-'.$regency['context']['name'];
            $districtPath = $regencyPath.DIRECTORY_SEPARATOR.$district['id'].'-'.$district['name'];
            $data['context'] = $district;
            $data['child'] = loadIndex($districtPath, INDEX_MAIN);
            $data['kb'] = loadData($districtPath, DISTRICT_KB);
            $data['tk'] = loadData($districtPath, DISTRICT_TK);
            $data['sd'] = loadData($districtPath, DISTRICT_SD);
            $data['smp'] = loadData($districtPath, DISTRICT_SMP);
            $data['sma'] = loadData($districtPath, DISTRICT_SMA);
            $data['smk'] = loadData($districtPath, DISTRICT_SMK);
            $data['slb'] = loadData($districtPath, DISTRICT_SLB);
            return $data;
        }
    }
    return null;
}