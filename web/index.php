<?php
include_once "config.php";
include_once "repo.php";

$id = "";
$data = null;
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

?>
<!DOCTYPE html>
<html>

<head>
    <?php if ($data['context'] == null) : ?>
        <title>Index Provinsi</title>
    <?php else : ?>
        <title>Index of <?= $data['context']['id'] . '-' . $data['context']['name'] ?></title>
    <?php endif; ?>
    <!--Import Google Icon Font-->
    <link href="https://fonts.googleapis.com/icon?family=Material+Icons" rel="stylesheet">
    <!--Import materialize.css-->
    <link rel="stylesheet" type="text/css" media="screen,projection" href="https://cdn.jsdelivr.net/npm/@materializecss/materialize@1.1.0/dist/css/materialize.min.css">

    <!--Let browser know website is optimized for mobile-->
    <meta name="viewport" content="width=device-width, initial-scale=1.0" />
</head>

<body>
    <div class="container">
        <?php if ($data['context'] == null) : ?>
            <h1>Index Provinsi<?= substr("3510", 0, 2); ?></h1>
        <?php else : ?>
            <h1>Index of <?= $data['context']['name'] ?></h1>
            <h3>[<?= $data['context']['id'] ?>]</h3>
            <a href="/">Home</a>
        <?php endif; ?>
        <?php if (strlen($id) == 4) : ?>
            <ul id="tabs-swipe-demo" class="tabs">
                <li class="tab col s3"><a class="active" href="#regency-district">Kecamatan</a></li>
                <li class="tab col s3"><a href="#regency-ponpes">Pondok Pesantren</a></li>
            </ul>
            <div id="regency-district" class="col s12">
                <div class="collection">
                    <?php foreach ($data['child'] as $district) : ?>
                        <a href="?id=<?= $district['id'] ?>" class="collection-item">[<?= $district['id'] ?>] <?= $district['name'] ?></a>
                    <?php endforeach; ?>
                </div>
            </div>
            <div id="regency-ponpes" class="col s12">
                <ul class="collection ">
                    <?php if (count($data['ponpes']) != 0) : ?>
                        <?php foreach ($data['ponpes'] as $ponpes) : ?>
                            <?php
                            if ($ponpes['district']['id'] == "") {
                                $status = "explore_off";
                                $color = "red-text";
                            } else {
                                $status = "explore";
                                $color = "green-text";
                            }
                            ?>
                            <li class="collection-item">
                                <div>
                                    [<?= $ponpes['id'] ?>] <?= $ponpes['nama'] ?>
                                    <a href="#!" class="secondary-content">
                                        <i class="material-icons <?= $color ?>"><?= $status ?></i>
                                    </a>
                                </div>
                            </li>
                        <?php endforeach; ?>
                    <?php else : ?>
                        <h5 class="center-align">Data pondok pesantren belum ter-indeks</h5>
                    <?php endif; ?>
                </ul>
            </div>
        <?php else : ?>
            <ul id="tabs-swipe-demo" class="tabs">
                <li class="tab col s3"><a class="active" href="#district-village">Dusun</a></li>
                <li class="tab col s3"><a href="#district-kb">KB</a></li>
                <li class="tab col s3"><a href="#district-tk">TK</a></li>
                <li class="tab col s3"><a href="#district-sd">SD</a></li>
                <li class="tab col s3"><a href="#district-smp">SMP</a></li>
                <li class="tab col s3"><a href="#district-sma">SMA</a></li>
                <li class="tab col s3"><a href="#district-smk">SMK</a></li>
                <li class="tab col s3"><a href="#district-slb">SLB</a></li>
            </ul>
            <div id="district-village" class="col s12">
                <div class="collection">
                    <?php foreach ($data['child'] as  $village) : ?>
                        <a href="?id=<?= $village['id'] ?>" class="collection-item">[<?= $village['id'] ?>] <?= $village['name'] ?></a>
                    <?php endforeach; ?>
                </div>
            </div>
            <div id="district-kb" class="col s12">
                <div class="collection">
                    <?php foreach ($data['kb'] as  $kb) : ?>
                        <a href="<?= BASE_DAPODIK . $kb['sekolah_id'] ?>" class="collection-item" target="_blank"><?= $kb['nama'] ?></a>
                    <?php endforeach; ?>
                </div>
            </div>
            <div id="district-tk" class="col s12">
                <div class="collection">
                    <?php foreach ($data['tk'] as  $tk) : ?>
                        <a href="<?= BASE_DAPODIK . $tk['sekolah_id'] ?>" class="collection-item" target="_blank"><?= $tk['nama'] ?></a>
                    <?php endforeach; ?>
                </div>
            </div>
            <div id="district-sd" class="col s12">
                <div class="collection">
                    <?php foreach ($data['sd'] as  $sd) : ?>
                        <a href="<?= BASE_DAPODIK . $sd['sekolah_id'] ?>" class="collection-item" target="_blank"><?= $sd['nama'] ?></a>
                    <?php endforeach; ?>
                </div>
            </div>
            <div id="district-smp" class="col s12">
                <div class="collection">
                    <?php foreach ($data['smp'] as  $smp) : ?>
                        <a href="<?= BASE_DAPODIK . $smp['sekolah_id'] ?>" class="collection-item" target="_blank"><?= $smp['nama'] ?></a>
                    <?php endforeach; ?>
                </div>
            </div>
            <div id="district-sma" class="col s12">
                <div class="collection">
                    <?php foreach ($data['sma'] as  $sma) : ?>
                        <a href="<?= BASE_DAPODIK . $sma['sekolah_id'] ?>" class="collection-item" target="_blank"><?= $sma['nama'] ?></a>
                    <?php endforeach; ?>
                </div>
            </div>
            <div id="district-smk" class="col s12">
                <div class="collection">
                    <?php foreach ($data['smk'] as  $smk) : ?>
                        <a href="<?= BASE_DAPODIK . $smk['sekolah_id'] ?>" class="collection-item" target="_blank"><?= $smk['nama'] ?></a>
                    <?php endforeach; ?>
                </div>
            </div>
            <div id="district-slb" class="col s12">
                <div class="collection">
                    <?php foreach ($data['slb'] as  $slb) : ?>
                        <a href="<?= BASE_DAPODIK . $slb['sekolah_id'] ?>" class="collection-item" target="_blank"><?= $slb['nama'] ?></a>
                    <?php endforeach; ?>
                </div>
            </div>
        <?php endif; ?>
    </div>

    <!--JavaScript at end of body for optimized loading-->
    <script src="https://cdn.jsdelivr.net/npm/@materializecss/materialize@1.1.0/dist/js/materialize.min.js"></script>
    <script>
        const tabs = document.querySelectorAll('.tabs');
        M.Tabs.init(tabs);
    </script>
</body>

</html>