<?php
const RootDir = ".." . DIRECTORY_SEPARATOR . "output";
const MainIndex = "index.json";
const MainIndexPath = RootDir . DIRECTORY_SEPARATOR . MainIndex;

const BASE_DAPODIK = "https://sekolah.data.kemdikbud.go.id/index.php/Chome/profil/";

const INDEX_MAIN = 0;
const INDEX_PONPES = 1;
const INDEX_UNIVERSITY = 2;

function loadIndex(string $path, int $type)
{
    $target = "";
    $json = [];
    switch ($type) {
        case INDEX_MAIN:
            $target = $path . DIRECTORY_SEPARATOR . MainIndex;
            break;
        case INDEX_PONPES:
            $target = $path . DIRECTORY_SEPARATOR . "ponpes-" . MainIndex;
            break;
        case INDEX_UNIVERSITY:
            $target = $path . DIRECTORY_SEPARATOR . "pt-" . MainIndex;
            break;
    }
    if (file_exists($target)) {
        $json = json_decode(file_get_contents($target), true);
    }
    return $json;
}

const DISTRICT_KB = 0;
const DISTRICT_TK = 1;
const DISTRICT_SD = 2;
const DISTRICT_SMP = 3;
const DISTRICT_SMA = 4;
const DISTRICT_SMK = 5;
const DISTRICT_SLB = 6;

function loadData(string $path, int $type)
{
    $json = [];
    $split = explode(DIRECTORY_SEPARATOR, $path);
    $target = $path . DIRECTORY_SEPARATOR . $split[count($split) - 1];
    switch ($type) {
        case DISTRICT_KB:
            $target = $target . "-kb.json";
            break;
        case DISTRICT_TK:
            $target = $target . "-tk.json";
            break;
        case DISTRICT_SD:
            $target = $target . "-sd.json";
            break;
        case DISTRICT_SMP:
            $target = $target . "-smp.json";
            break;
        case DISTRICT_SMA:
            $target = $target . "-sma.json";
            break;
        case DISTRICT_SMK:
            $target = $target . "-smk.json";
            break;
        case DISTRICT_SLB:
            $target = $target . "-slb.json";
            break;
    }
    if (file_exists($target)) {
        $json = json_decode(file_get_contents($target), true);
    }
    return $json;
}
